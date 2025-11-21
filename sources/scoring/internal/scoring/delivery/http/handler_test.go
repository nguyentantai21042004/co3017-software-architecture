package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"scoring/internal/scoring"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// Mock UseCase
type MockUseCase struct {
	mock.Mock
}

func (m *MockUseCase) SubmitAnswer(ctx context.Context, input scoring.SubmitInput) (scoring.SubmitOutput, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(scoring.SubmitOutput), args.Error(1)
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

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.New()
}

func TestSubmitAnswer_Success(t *testing.T) {
	mockLogger := new(MockLogger)
	mockUC := new(MockUseCase)

	mockLogger.On("Errorf", mock.Anything, mock.Anything, mock.Anything).Return()

	expectedOutput := scoring.SubmitOutput{
		Correct:  true,
		Score:    100,
		Feedback: "Correct! Well done.",
	}
	mockUC.On("SubmitAnswer", mock.Anything, mock.MatchedBy(func(input scoring.SubmitInput) bool {
		return input.UserID == "user123" && input.QuestionID == 1 && input.Answer == "4"
	})).Return(expectedOutput, nil)

	handler := New(mockLogger, mockUC)
	router := setupRouter()
	router.POST("/submit", handler.SubmitAnswer)

	reqBody := SubmitRequest{
		UserID:     "user123",
		QuestionID: 1,
		Answer:     "4",
	}
	body, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", "/submit", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, float64(0), response["error_code"])
	assert.Equal(t, "Success", response["message"])

	data := response["data"].(map[string]interface{})
	assert.Equal(t, true, data["correct"])
	assert.Equal(t, float64(100), data["score"])
	assert.Equal(t, "Correct! Well done.", data["feedback"])

	mockUC.AssertExpectations(t)
}

func TestSubmitAnswer_InvalidJSON(t *testing.T) {
	mockLogger := new(MockLogger)
	mockUC := new(MockUseCase)

	mockLogger.On("Errorf", mock.Anything, mock.Anything, mock.Anything).Return()

	handler := New(mockLogger, mockUC)
	router := setupRouter()
	router.POST("/submit", handler.SubmitAnswer)

	req, _ := http.NewRequest("POST", "/submit", bytes.NewBufferString("{invalid json"))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, float64(400), response["error_code"])
	assert.Contains(t, response["message"], "failed to bind request")
}

func TestSubmitAnswer_EmptyUserID(t *testing.T) {
	mockLogger := new(MockLogger)
	mockUC := new(MockUseCase)

	mockLogger.On("Errorf", mock.Anything, mock.Anything, mock.Anything).Return()

	handler := New(mockLogger, mockUC)
	router := setupRouter()
	router.POST("/submit", handler.SubmitAnswer)

	reqBody := SubmitRequest{
		UserID:     "",
		QuestionID: 1,
		Answer:     "4",
	}
	body, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", "/submit", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, float64(400), response["error_code"])
	assert.Contains(t, response["message"], "user_id")
}

func TestSubmitAnswer_InvalidQuestionID(t *testing.T) {
	mockLogger := new(MockLogger)
	mockUC := new(MockUseCase)

	mockLogger.On("Errorf", mock.Anything, mock.Anything, mock.Anything).Return()

	handler := New(mockLogger, mockUC)
	router := setupRouter()
	router.POST("/submit", handler.SubmitAnswer)

	reqBody := SubmitRequest{
		UserID:     "user123",
		QuestionID: 0,
		Answer:     "4",
	}
	body, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", "/submit", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, float64(400), response["error_code"])
	assert.Contains(t, response["message"], "question_id")
}

func TestSubmitAnswer_EmptyAnswer(t *testing.T) {
	mockLogger := new(MockLogger)
	mockUC := new(MockUseCase)

	mockLogger.On("Errorf", mock.Anything, mock.Anything, mock.Anything).Return()

	handler := New(mockLogger, mockUC)
	router := setupRouter()
	router.POST("/submit", handler.SubmitAnswer)

	reqBody := SubmitRequest{
		UserID:     "user123",
		QuestionID: 1,
		Answer:     "",
	}
	body, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", "/submit", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, float64(400), response["error_code"])
	assert.Contains(t, response["message"], "answer")
}

func TestSubmitAnswer_UseCaseError(t *testing.T) {
	mockLogger := new(MockLogger)
	mockUC := new(MockUseCase)

	mockLogger.On("Errorf", mock.Anything, mock.Anything, mock.Anything).Return()

	expectedErr := errors.New("usecase error")
	mockUC.On("SubmitAnswer", mock.Anything, mock.Anything).Return(scoring.SubmitOutput{}, expectedErr)

	handler := New(mockLogger, mockUC)
	router := setupRouter()
	router.POST("/submit", handler.SubmitAnswer)

	reqBody := SubmitRequest{
		UserID:     "user123",
		QuestionID: 1,
		Answer:     "4",
	}
	body, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", "/submit", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, float64(400), response["error_code"])

	mockUC.AssertExpectations(t)
}

func TestHealth(t *testing.T) {
	mockLogger := new(MockLogger)
	mockUC := new(MockUseCase)

	handler := New(mockLogger, mockUC)
	router := setupRouter()
	router.GET("/health", handler.Health)

	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, float64(0), response["error_code"])
	assert.Equal(t, "Healthy", response["message"])

	data := response["data"].(map[string]interface{})
	assert.Equal(t, "healthy", data["status"])
	assert.Equal(t, "scoring", data["service"])
}

func TestNew(t *testing.T) {
	mockLogger := new(MockLogger)
	mockUC := new(MockUseCase)

	handler := New(mockLogger, mockUC)

	assert.NotNil(t, handler)
}
