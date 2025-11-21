package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"scoring/internal/model"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockScoringService
type MockScoringService struct {
	mock.Mock
}

func (m *MockScoringService) SubmitAnswer(req *model.SubmitRequest) (*model.SubmitResponse, error) {
	args := m.Called(req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.SubmitResponse), args.Error(1)
}

func TestSubmitAnswer_Success(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	mockService := new(MockScoringService)
	handler := NewScoringHandler(mockService)
	router := gin.Default()
	router.POST("/submit", handler.SubmitAnswer)

	reqBody := &model.SubmitRequest{
		UserID:     "user1",
		QuestionID: 1,
		Answer:     "A",
	}
	jsonBody, _ := json.Marshal(reqBody)

	expectedResp := &model.SubmitResponse{
		Correct:  true,
		Score:    100,
		Feedback: "Correct",
	}

	mockService.On("SubmitAnswer", reqBody).Return(expectedResp, nil)

	// Execute
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/submit", bytes.NewBuffer(jsonBody))
	router.ServeHTTP(w, req)

	// Verify
	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestSubmitAnswer_BadRequest(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	mockService := new(MockScoringService)
	handler := NewScoringHandler(mockService)
	router := gin.Default()
	router.POST("/submit", handler.SubmitAnswer)

	// Invalid JSON
	reqBody := []byte(`{"invalid": "json"}`)

	// Execute
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/submit", bytes.NewBuffer(reqBody))
	router.ServeHTTP(w, req)

	// Verify
	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockService.AssertNotCalled(t, "SubmitAnswer")
}

func TestSubmitAnswer_ServiceError(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	mockService := new(MockScoringService)
	handler := NewScoringHandler(mockService)
	router := gin.Default()
	router.POST("/submit", handler.SubmitAnswer)

	reqBody := &model.SubmitRequest{
		UserID:     "user1",
		QuestionID: 1,
		Answer:     "A",
	}
	jsonBody, _ := json.Marshal(reqBody)

	mockService.On("SubmitAnswer", reqBody).Return(nil, errors.New("service error"))

	// Execute
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/submit", bytes.NewBuffer(jsonBody))
	router.ServeHTTP(w, req)

	// Verify
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockService.AssertExpectations(t)
}
