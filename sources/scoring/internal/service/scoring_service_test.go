package service

import (
	"errors"
	"scoring/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockSubmissionRepository
type MockSubmissionRepository struct {
	mock.Mock
}

func (m *MockSubmissionRepository) Create(submission *model.Submission) error {
	args := m.Called(submission)
	return args.Error(0)
}

// MockEventPublisher
type MockEventPublisher struct {
	mock.Mock
	Done chan bool
}

func (m *MockEventPublisher) PublishSubmissionEvent(event interface{}) error {
	args := m.Called(event)
	if m.Done != nil {
		m.Done <- true
	}
	return args.Error(0)
}

func (m *MockEventPublisher) Close() error {
	args := m.Called()
	return args.Error(0)
}

// MockContentClient
type MockContentClient struct {
	mock.Mock
}

func (m *MockContentClient) FetchQuestion(questionID int64) (string, string, error) {
	args := m.Called(questionID)
	return args.String(0), args.String(1), args.Error(2)
}

func TestSubmitAnswer_Correct(t *testing.T) {
	// Setup
	mockRepo := new(MockSubmissionRepository)
	mockPub := new(MockEventPublisher)
	mockPub.Done = make(chan bool, 1) // Buffer to prevent blocking
	mockClient := new(MockContentClient)
	service := NewScoringService(mockRepo, mockPub, mockClient)

	req := &model.SubmitRequest{
		UserID:     "user1",
		QuestionID: 1,
		Answer:     "A",
	}

	// Expectations
	mockClient.On("FetchQuestion", int64(1)).Return("A", "math", nil)
	mockRepo.On("Create", mock.AnythingOfType("*model.Submission")).Return(nil)
	mockPub.On("PublishSubmissionEvent", mock.AnythingOfType("model.SubmissionEvent")).Return(nil)

	// Execute
	resp, err := service.SubmitAnswer(req)

	// Verify
	assert.NoError(t, err)
	assert.True(t, resp.Correct)
	assert.Equal(t, 100, resp.Score)

	// Wait for async
	<-mockPub.Done

	mockClient.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
	mockPub.AssertExpectations(t)
}

func TestSubmitAnswer_Incorrect(t *testing.T) {
	// Setup
	mockRepo := new(MockSubmissionRepository)
	mockPub := new(MockEventPublisher)
	mockPub.Done = make(chan bool, 1)
	mockClient := new(MockContentClient)
	service := NewScoringService(mockRepo, mockPub, mockClient)

	req := &model.SubmitRequest{
		UserID:     "user1",
		QuestionID: 1,
		Answer:     "B",
	}

	// Expectations
	mockClient.On("FetchQuestion", int64(1)).Return("A", "math", nil)
	mockRepo.On("Create", mock.AnythingOfType("*model.Submission")).Return(nil)
	mockPub.On("PublishSubmissionEvent", mock.AnythingOfType("model.SubmissionEvent")).Return(nil)

	// Execute
	resp, err := service.SubmitAnswer(req)

	// Verify
	assert.NoError(t, err)
	assert.False(t, resp.Correct)
	assert.Equal(t, 0, resp.Score)

	// Wait for async
	<-mockPub.Done

	mockClient.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
	mockPub.AssertExpectations(t)
}

func TestSubmitAnswer_ContentServiceError(t *testing.T) {
	// Setup
	mockRepo := new(MockSubmissionRepository)
	mockPub := new(MockEventPublisher)
	mockClient := new(MockContentClient)
	service := NewScoringService(mockRepo, mockPub, mockClient)

	req := &model.SubmitRequest{
		UserID:     "user1",
		QuestionID: 1,
		Answer:     "A",
	}

	// Expectations
	mockClient.On("FetchQuestion", int64(1)).Return("", "", errors.New("service down"))

	// Execute
	resp, err := service.SubmitAnswer(req)

	// Verify
	assert.Error(t, err)
	assert.Nil(t, resp)
	mockClient.AssertExpectations(t)
	mockRepo.AssertNotCalled(t, "Create")
	mockPub.AssertNotCalled(t, "PublishSubmissionEvent")
}
