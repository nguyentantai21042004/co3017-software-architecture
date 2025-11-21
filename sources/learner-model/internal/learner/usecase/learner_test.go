package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"learner-model-service/internal/learner"
	"learner-model-service/internal/learner/repository"
	"learner-model-service/internal/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRepository is a mock implementation of repository.Repository
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetByUserAndSkill(ctx context.Context, userID, skillTag string) (*model.SkillMastery, error) {
	args := m.Called(ctx, userID, skillTag)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.SkillMastery), args.Error(1)
}

func (m *MockRepository) CreateOrUpdate(ctx context.Context, mastery *model.SkillMastery) error {
	args := m.Called(ctx, mastery)
	return args.Error(0)
}

// MockLogger is a mock implementation of log.Logger
type MockLogger struct {
	mock.Mock
}

func (m *MockLogger) Debug(ctx context.Context, arg ...any) {
	m.Called(ctx, arg)
}

func (m *MockLogger) Debugf(ctx context.Context, format string, arg ...any) {
	m.Called(ctx, format, arg)
}

func (m *MockLogger) Info(ctx context.Context, arg ...any) {
	m.Called(ctx, arg)
}

func (m *MockLogger) Infof(ctx context.Context, format string, arg ...any) {
	m.Called(ctx, format, arg)
}

func (m *MockLogger) Warn(ctx context.Context, arg ...any) {
	m.Called(ctx, arg)
}

func (m *MockLogger) Warnf(ctx context.Context, format string, arg ...any) {
	m.Called(ctx, format, arg)
}

func (m *MockLogger) Error(ctx context.Context, arg ...any) {
	m.Called(ctx, arg)
}

func (m *MockLogger) Errorf(ctx context.Context, format string, arg ...any) {
	m.Called(ctx, format, arg)
}

func (m *MockLogger) Fatal(ctx context.Context, arg ...any) {
	m.Called(ctx, arg)
}

func (m *MockLogger) Fatalf(ctx context.Context, format string, arg ...any) {
	m.Called(ctx, format, arg)
}

func TestGetMastery_Success_ExistingMastery(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := new(MockRepository)
	mockLogger := new(MockLogger)

	uc := &usecase{
		l:    mockLogger,
		repo: mockRepo,
	}

	userID := "user123"
	skillTag := "math_algebra"
	expectedMastery := &model.SkillMastery{
		UserID:       userID,
		SkillTag:     skillTag,
		CurrentScore: 75,
		LastUpdated:  time.Now(),
	}

	input := learner.GetMasteryInput{
		UserID:   userID,
		SkillTag: skillTag,
	}

	// Mock expectations
	mockLogger.On("Infof", ctx, mock.Anything, mock.Anything).Return()
	mockRepo.On("GetByUserAndSkill", ctx, userID, skillTag).Return(expectedMastery, nil)

	// Act
	output, err := uc.GetMastery(ctx, input)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, userID, output.UserID)
	assert.Equal(t, skillTag, output.SkillTag)
	assert.Equal(t, 75, output.MasteryScore)
	mockRepo.AssertExpectations(t)
}

func TestGetMastery_Success_MasteryNotFound_ReturnsDefault(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := new(MockRepository)
	mockLogger := new(MockLogger)

	uc := &usecase{
		l:    mockLogger,
		repo: mockRepo,
	}

	userID := "user456"
	skillTag := "science_physics"

	input := learner.GetMasteryInput{
		UserID:   userID,
		SkillTag: skillTag,
	}

	// Mock expectations
	mockLogger.On("Infof", ctx, mock.Anything, mock.Anything).Return()
	mockRepo.On("GetByUserAndSkill", ctx, userID, skillTag).Return(nil, repository.ErrNotFound)

	// Act
	output, err := uc.GetMastery(ctx, input)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, userID, output.UserID)
	assert.Equal(t, skillTag, output.SkillTag)
	assert.Equal(t, 0, output.MasteryScore) // Default score
	mockRepo.AssertExpectations(t)
}

func TestGetMastery_Failure_RepositoryError(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := new(MockRepository)
	mockLogger := new(MockLogger)

	uc := &usecase{
		l:    mockLogger,
		repo: mockRepo,
	}

	userID := "user789"
	skillTag := "math_calculus"

	input := learner.GetMasteryInput{
		UserID:   userID,
		SkillTag: skillTag,
	}

	dbError := errors.New("database connection failed")

	// Mock expectations
	mockLogger.On("Infof", ctx, mock.Anything, mock.Anything).Return()
	mockLogger.On("Errorf", ctx, mock.Anything, mock.Anything).Return()
	mockRepo.On("GetByUserAndSkill", ctx, userID, skillTag).Return(nil, dbError)

	// Act
	output, err := uc.GetMastery(ctx, input)

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), ErrMsgGetMasteryFailed)
	assert.Equal(t, "", output.UserID)
	mockRepo.AssertExpectations(t)
}

func TestUpdateMasteryFromEvent_Success_NewMastery(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := new(MockRepository)
	mockLogger := new(MockLogger)

	uc := &usecase{
		l:    mockLogger,
		repo: mockRepo,
	}

	userID := "user123"
	skillTag := "math_algebra"
	scoreObtained := 80

	input := learner.UpdateMasteryInput{
		UserID:        userID,
		SkillTag:      skillTag,
		ScoreObtained: scoreObtained,
	}

	// Mock expectations - no existing mastery
	mockLogger.On("Infof", ctx, mock.Anything, mock.Anything).Return()
	mockRepo.On("GetByUserAndSkill", ctx, userID, skillTag).Return(nil, repository.ErrNotFound)
	mockRepo.On("CreateOrUpdate", ctx, mock.MatchedBy(func(m *model.SkillMastery) bool {
		// New score should be (0 + 80) / 2 = 40
		return m.UserID == userID && m.SkillTag == skillTag && m.CurrentScore == 40
	})).Return(nil)

	// Act
	err := uc.UpdateMasteryFromEvent(ctx, input)

	// Assert
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateMasteryFromEvent_Success_ExistingMastery(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := new(MockRepository)
	mockLogger := new(MockLogger)

	uc := &usecase{
		l:    mockLogger,
		repo: mockRepo,
	}

	userID := "user456"
	skillTag := "science_chemistry"
	oldScore := 60
	scoreObtained := 80

	input := learner.UpdateMasteryInput{
		UserID:        userID,
		SkillTag:      skillTag,
		ScoreObtained: scoreObtained,
	}

	existingMastery := &model.SkillMastery{
		UserID:       userID,
		SkillTag:     skillTag,
		CurrentScore: oldScore,
		LastUpdated:  time.Now(),
	}

	// Mock expectations
	mockLogger.On("Infof", ctx, mock.Anything, mock.Anything).Return()
	mockRepo.On("GetByUserAndSkill", ctx, userID, skillTag).Return(existingMastery, nil)
	mockRepo.On("CreateOrUpdate", ctx, mock.MatchedBy(func(m *model.SkillMastery) bool {
		// New score should be (60 + 80) / 2 = 70
		return m.UserID == userID && m.SkillTag == skillTag && m.CurrentScore == 70
	})).Return(nil)

	// Act
	err := uc.UpdateMasteryFromEvent(ctx, input)

	// Assert
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateMasteryFromEvent_ScoreBoundaries(t *testing.T) {
	tests := []struct {
		name           string
		oldScore       int
		scoreObtained  int
		expectedScore  int
	}{
		{
			name:          "Score below 0 should be clamped to 0",
			oldScore:      0,
			scoreObtained: -50,
			expectedScore: 0,
		},
		{
			name:          "Score above 100 should be clamped to 100",
			oldScore:      100,
			scoreObtained: 100,
			expectedScore: 100,
		},
		{
			name:          "Normal score calculation",
			oldScore:      50,
			scoreObtained: 70,
			expectedScore: 60,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			ctx := context.Background()
			mockRepo := new(MockRepository)
			mockLogger := new(MockLogger)

			uc := &usecase{
				l:    mockLogger,
				repo: mockRepo,
			}

			userID := "user_test"
			skillTag := "test_skill"

			input := learner.UpdateMasteryInput{
				UserID:        userID,
				SkillTag:      skillTag,
				ScoreObtained: tt.scoreObtained,
			}

			var existingMastery *model.SkillMastery
			if tt.oldScore > 0 {
				existingMastery = &model.SkillMastery{
					UserID:       userID,
					SkillTag:     skillTag,
					CurrentScore: tt.oldScore,
					LastUpdated:  time.Now(),
				}
			}

			// Mock expectations
			mockLogger.On("Infof", ctx, mock.Anything, mock.Anything).Return()
			if existingMastery != nil {
				mockRepo.On("GetByUserAndSkill", ctx, userID, skillTag).Return(existingMastery, nil)
			} else {
				mockRepo.On("GetByUserAndSkill", ctx, userID, skillTag).Return(nil, repository.ErrNotFound)
			}
			mockRepo.On("CreateOrUpdate", ctx, mock.MatchedBy(func(m *model.SkillMastery) bool {
				return m.UserID == userID && m.SkillTag == skillTag && m.CurrentScore == tt.expectedScore
			})).Return(nil)

			// Act
			err := uc.UpdateMasteryFromEvent(ctx, input)

			// Assert
			assert.NoError(t, err)
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestUpdateMasteryFromEvent_Failure_GetMasteryError(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := new(MockRepository)
	mockLogger := new(MockLogger)

	uc := &usecase{
		l:    mockLogger,
		repo: mockRepo,
	}

	userID := "user789"
	skillTag := "math_geometry"

	input := learner.UpdateMasteryInput{
		UserID:        userID,
		SkillTag:      skillTag,
		ScoreObtained: 75,
	}

	dbError := errors.New("database connection failed")

	// Mock expectations
	mockLogger.On("Infof", ctx, mock.Anything, mock.Anything).Return()
	mockLogger.On("Errorf", ctx, mock.Anything, mock.Anything).Return()
	mockRepo.On("GetByUserAndSkill", ctx, userID, skillTag).Return(nil, dbError)

	// Act
	err := uc.UpdateMasteryFromEvent(ctx, input)

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), ErrMsgUpdateMasteryFailed)
	mockRepo.AssertExpectations(t)
	mockRepo.AssertNotCalled(t, "CreateOrUpdate", mock.Anything, mock.Anything)
}

func TestUpdateMasteryFromEvent_Failure_CreateOrUpdateError(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := new(MockRepository)
	mockLogger := new(MockLogger)

	uc := &usecase{
		l:    mockLogger,
		repo: mockRepo,
	}

	userID := "user101"
	skillTag := "english_grammar"

	input := learner.UpdateMasteryInput{
		UserID:        userID,
		SkillTag:      skillTag,
		ScoreObtained: 90,
	}

	dbError := errors.New("failed to insert record")

	// Mock expectations
	mockLogger.On("Infof", ctx, mock.Anything, mock.Anything).Return()
	mockLogger.On("Errorf", ctx, mock.Anything, mock.Anything).Return()
	mockRepo.On("GetByUserAndSkill", ctx, userID, skillTag).Return(nil, repository.ErrNotFound)
	mockRepo.On("CreateOrUpdate", ctx, mock.Anything).Return(dbError)

	// Act
	err := uc.UpdateMasteryFromEvent(ctx, input)

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), ErrMsgUpdateMasteryFailed)
	mockRepo.AssertExpectations(t)
}
