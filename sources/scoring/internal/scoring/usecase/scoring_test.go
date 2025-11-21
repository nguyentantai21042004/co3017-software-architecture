package usecase

import (
	"context"
	"errors"
	"testing"

	"scoring/internal/model"
	"scoring/internal/scoring"
	"scoring/pkg/curl"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// Mock Repository
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Create(submission *model.Submission) error {
	args := m.Called(submission)
	return args.Error(0)
}

// Mock Publisher
type MockPublisher struct {
	mock.Mock
}

func (m *MockPublisher) PublishSubmissionEvent(event interface{}) error {
	args := m.Called(event)
	return args.Error(0)
}

func (m *MockPublisher) Close() error {
	args := m.Called()
	return args.Error(0)
}

// Mock Content Service Client
type MockContentServiceClient struct {
	mock.Mock
}

func (m *MockContentServiceClient) GetQuestion(ctx context.Context, questionID int64) (*curl.ContentQuestionResponse, error) {
	args := m.Called(ctx, questionID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*curl.ContentQuestionResponse), args.Error(1)
}

// Mock Logger
type MockLogger struct {
	mock.Mock
}

func (m *MockLogger) Debug(ctx context.Context, arg ...any) {
	m.Called(ctx, arg)
}

func (m *MockLogger) Debugf(ctx context.Context, template string, arg ...any) {
	m.Called(ctx, template, arg)
}

func (m *MockLogger) Info(ctx context.Context, arg ...any) {
	m.Called(ctx, arg)
}

func (m *MockLogger) Infof(ctx context.Context, template string, arg ...any) {
	m.Called(ctx, template, arg)
}

func (m *MockLogger) Warn(ctx context.Context, arg ...any) {
	m.Called(ctx, arg)
}

func (m *MockLogger) Warnf(ctx context.Context, template string, arg ...any) {
	m.Called(ctx, template, arg)
}

func (m *MockLogger) Error(ctx context.Context, arg ...any) {
	m.Called(ctx, arg)
}

func (m *MockLogger) Errorf(ctx context.Context, template string, arg ...any) {
	m.Called(ctx, template, arg)
}

func (m *MockLogger) Fatal(ctx context.Context, arg ...any) {
	m.Called(ctx, arg)
}

func (m *MockLogger) Fatalf(ctx context.Context, template string, arg ...any) {
	m.Called(ctx, template, arg)
}

func TestSubmitAnswer_CorrectAnswer(t *testing.T) {
	// Setup mocks
	mockLogger := new(MockLogger)
	mockRepo := new(MockRepository)
	mockPublisher := new(MockPublisher)
	mockContentClient := new(MockContentServiceClient)

	mockLogger.On("Infof", mock.Anything, mock.Anything, mock.Anything).Return()
	mockLogger.On("Errorf", mock.Anything, mock.Anything, mock.Anything).Return()

	// Mock content service response
	contentResp := &curl.ContentQuestionResponse{
		Success: true,
		Message: "Success",
		Data: struct {
			ID            int64  `json:"id"`
			Content       string `json:"content"`
			CorrectAnswer string `json:"correct_answer"`
			SkillTag      string `json:"skill_tag"`
			IsRemedial    bool   `json:"is_remedial"`
		}{
			ID:            1,
			Content:       "What is 2+2?",
			CorrectAnswer: "4",
			SkillTag:      "math",
			IsRemedial:    false,
		},
	}
	mockContentClient.On("GetQuestion", mock.Anything, int64(1)).Return(contentResp, nil)

	// Mock repository
	mockRepo.On("Create", mock.MatchedBy(func(s *model.Submission) bool {
		return s.UserID == "user123" && s.QuestionID == 1 && s.ScoreAwarded == 100
	})).Return(nil)

	// Mock publisher (async call, may or may not be called immediately)
	mockPublisher.On("PublishSubmissionEvent", mock.Anything).Return(nil).Maybe()

	// Create usecase
	uc := New(mockLogger, mockRepo, mockPublisher, mockContentClient)

	// Test
	input := scoring.SubmitInput{
		UserID:     "user123",
		QuestionID: 1,
		Answer:     "4",
	}

	result, err := uc.SubmitAnswer(context.Background(), input)

	// Assertions
	require.NoError(t, err)
	assert.True(t, result.Correct)
	assert.Equal(t, 100, result.Score)
	assert.Equal(t, "Correct! Well done.", result.Feedback)

	mockContentClient.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestSubmitAnswer_IncorrectAnswer(t *testing.T) {
	mockLogger := new(MockLogger)
	mockRepo := new(MockRepository)
	mockPublisher := new(MockPublisher)
	mockContentClient := new(MockContentServiceClient)

	mockLogger.On("Infof", mock.Anything, mock.Anything, mock.Anything).Return()
	mockLogger.On("Errorf", mock.Anything, mock.Anything, mock.Anything).Return()

	contentResp := &curl.ContentQuestionResponse{
		Success: true,
		Message: "Success",
		Data: struct {
			ID            int64  `json:"id"`
			Content       string `json:"content"`
			CorrectAnswer string `json:"correct_answer"`
			SkillTag      string `json:"skill_tag"`
			IsRemedial    bool   `json:"is_remedial"`
		}{
			ID:            1,
			Content:       "What is 2+2?",
			CorrectAnswer: "4",
			SkillTag:      "math",
		},
	}
	mockContentClient.On("GetQuestion", mock.Anything, int64(1)).Return(contentResp, nil)

	mockRepo.On("Create", mock.MatchedBy(func(s *model.Submission) bool {
		return s.UserID == "user123" && s.ScoreAwarded == 0 && !s.IsPassed
	})).Return(nil)

	mockPublisher.On("PublishSubmissionEvent", mock.Anything).Return(nil).Maybe()

	uc := New(mockLogger, mockRepo, mockPublisher, mockContentClient)

	input := scoring.SubmitInput{
		UserID:     "user123",
		QuestionID: 1,
		Answer:     "5", // Wrong answer
	}

	result, err := uc.SubmitAnswer(context.Background(), input)

	require.NoError(t, err)
	assert.False(t, result.Correct)
	assert.Equal(t, 0, result.Score)
	assert.Equal(t, "Incorrect answer. Please try again!", result.Feedback)

	mockContentClient.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestSubmitAnswer_FetchQuestionError(t *testing.T) {
	mockLogger := new(MockLogger)
	mockRepo := new(MockRepository)
	mockPublisher := new(MockPublisher)
	mockContentClient := new(MockContentServiceClient)

	mockLogger.On("Infof", mock.Anything, mock.Anything, mock.Anything).Return()
	mockLogger.On("Errorf", mock.Anything, mock.Anything, mock.Anything).Return()

	expectedErr := errors.New("service unavailable")
	mockContentClient.On("GetQuestion", mock.Anything, int64(1)).Return(nil, expectedErr)

	uc := New(mockLogger, mockRepo, mockPublisher, mockContentClient)

	input := scoring.SubmitInput{
		UserID:     "user123",
		QuestionID: 1,
		Answer:     "4",
	}

	result, err := uc.SubmitAnswer(context.Background(), input)

	require.Error(t, err)
	assert.Contains(t, err.Error(), ErrMsgFetchQuestionFailed)
	assert.Equal(t, scoring.SubmitOutput{}, result)

	mockContentClient.AssertExpectations(t)
}

func TestSubmitAnswer_SaveSubmissionError(t *testing.T) {
	mockLogger := new(MockLogger)
	mockRepo := new(MockRepository)
	mockPublisher := new(MockPublisher)
	mockContentClient := new(MockContentServiceClient)

	mockLogger.On("Infof", mock.Anything, mock.Anything, mock.Anything).Return()
	mockLogger.On("Errorf", mock.Anything, mock.Anything, mock.Anything).Return()

	contentResp := &curl.ContentQuestionResponse{
		Success: true,
		Message: "Success",
		Data: struct {
			ID            int64  `json:"id"`
			Content       string `json:"content"`
			CorrectAnswer string `json:"correct_answer"`
			SkillTag      string `json:"skill_tag"`
			IsRemedial    bool   `json:"is_remedial"`
		}{
			ID:            1,
			CorrectAnswer: "4",
			SkillTag:      "math",
		},
	}
	mockContentClient.On("GetQuestion", mock.Anything, int64(1)).Return(contentResp, nil)

	expectedErr := errors.New("database error")
	mockRepo.On("Create", mock.Anything).Return(expectedErr)

	uc := New(mockLogger, mockRepo, mockPublisher, mockContentClient)

	input := scoring.SubmitInput{
		UserID:     "user123",
		QuestionID: 1,
		Answer:     "4",
	}

	result, err := uc.SubmitAnswer(context.Background(), input)

	require.Error(t, err)
	assert.Contains(t, err.Error(), ErrMsgSaveSubmissionFailed)
	assert.Equal(t, scoring.SubmitOutput{}, result)

	mockContentClient.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestNew(t *testing.T) {
	mockLogger := new(MockLogger)
	mockRepo := new(MockRepository)
	mockPublisher := new(MockPublisher)
	mockContentClient := new(MockContentServiceClient)

	uc := New(mockLogger, mockRepo, mockPublisher, mockContentClient)

	assert.NotNil(t, uc)
}
