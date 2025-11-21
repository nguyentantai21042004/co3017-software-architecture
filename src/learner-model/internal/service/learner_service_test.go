package service

import (
	"errors"
	"learner-model-service/internal/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockMasteryRepository
type MockMasteryRepository struct {
	mock.Mock
}

func (m *MockMasteryRepository) GetByUserAndSkill(userID, skillTag string) (*model.SkillMastery, error) {
	args := m.Called(userID, skillTag)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.SkillMastery), args.Error(1)
}

func (m *MockMasteryRepository) UpdateScore(userID, skillTag string, newScore int) error {
	args := m.Called(userID, skillTag, newScore)
	return args.Error(0)
}

func (m *MockMasteryRepository) CreateOrUpdate(mastery *model.SkillMastery) error {
	args := m.Called(mastery)
	return args.Error(0)
}

func TestGetMastery_Exists(t *testing.T) {
	// Setup
	mockRepo := new(MockMasteryRepository)
	service := NewLearnerService(mockRepo)

	existingMastery := &model.SkillMastery{
		UserID:       "user1",
		SkillTag:     "math",
		CurrentScore: 75,
		LastUpdated:  time.Now(),
	}

	mockRepo.On("GetByUserAndSkill", "user1", "math").Return(existingMastery, nil)

	// Execute
	resp, err := service.GetMastery("user1", "math")

	// Verify
	assert.NoError(t, err)
	assert.Equal(t, "user1", resp.UserID)
	assert.Equal(t, "math", resp.SkillTag)
	assert.Equal(t, 75, resp.MasteryScore)
	mockRepo.AssertExpectations(t)
}

func TestGetMastery_NotExists(t *testing.T) {
	// Setup
	mockRepo := new(MockMasteryRepository)
	service := NewLearnerService(mockRepo)

	mockRepo.On("GetByUserAndSkill", "user1", "math").Return(nil, nil)

	// Execute
	resp, err := service.GetMastery("user1", "math")

	// Verify
	assert.NoError(t, err)
	assert.Equal(t, "user1", resp.UserID)
	assert.Equal(t, "math", resp.SkillTag)
	assert.Equal(t, 0, resp.MasteryScore) // Default score
	mockRepo.AssertExpectations(t)
}

func TestUpdateMasteryFromEvent_NewUser(t *testing.T) {
	// Setup
	mockRepo := new(MockMasteryRepository)
	service := NewLearnerService(mockRepo)

	event := &model.SubmissionEvent{
		Event:         "SubmissionCompleted",
		UserID:        "user1",
		SkillTag:      "math",
		ScoreObtained: 80,
	}

	// No existing mastery
	mockRepo.On("GetByUserAndSkill", "user1", "math").Return(nil, nil)
	mockRepo.On("CreateOrUpdate", mock.MatchedBy(func(m *model.SkillMastery) bool {
		return m.UserID == "user1" && m.SkillTag == "math" && m.CurrentScore == 40 // (0 + 80) / 2
	})).Return(nil)

	// Execute
	err := service.UpdateMasteryFromEvent(event)

	// Verify
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateMasteryFromEvent_ExistingUser(t *testing.T) {
	// Setup
	mockRepo := new(MockMasteryRepository)
	service := NewLearnerService(mockRepo)

	event := &model.SubmissionEvent{
		Event:         "SubmissionCompleted",
		UserID:        "user1",
		SkillTag:      "math",
		ScoreObtained: 100,
	}

	existingMastery := &model.SkillMastery{
		UserID:       "user1",
		SkillTag:     "math",
		CurrentScore: 60,
	}

	mockRepo.On("GetByUserAndSkill", "user1", "math").Return(existingMastery, nil)
	mockRepo.On("CreateOrUpdate", mock.MatchedBy(func(m *model.SkillMastery) bool {
		return m.UserID == "user1" && m.SkillTag == "math" && m.CurrentScore == 80 // (60 + 100) / 2
	})).Return(nil)

	// Execute
	err := service.UpdateMasteryFromEvent(event)

	// Verify
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateMasteryFromEvent_ScoreBounds(t *testing.T) {
	// Setup
	mockRepo := new(MockMasteryRepository)
	service := NewLearnerService(mockRepo)

	// Test upper bound
	event := &model.SubmissionEvent{
		UserID:        "user1",
		SkillTag:      "math",
		ScoreObtained: 150, // Over 100
	}

	existingMastery := &model.SkillMastery{
		UserID:       "user1",
		SkillTag:     "math",
		CurrentScore: 100,
	}

	mockRepo.On("GetByUserAndSkill", "user1", "math").Return(existingMastery, nil)
	mockRepo.On("CreateOrUpdate", mock.MatchedBy(func(m *model.SkillMastery) bool {
		return m.CurrentScore == 100 // Capped at 100
	})).Return(nil)

	// Execute
	err := service.UpdateMasteryFromEvent(event)

	// Verify
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateMasteryFromEvent_RepositoryError(t *testing.T) {
	// Setup
	mockRepo := new(MockMasteryRepository)
	service := NewLearnerService(mockRepo)

	event := &model.SubmissionEvent{
		UserID:        "user1",
		SkillTag:      "math",
		ScoreObtained: 80,
	}

	mockRepo.On("GetByUserAndSkill", "user1", "math").Return(nil, errors.New("db error"))

	// Execute
	err := service.UpdateMasteryFromEvent(event)

	// Verify
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}
