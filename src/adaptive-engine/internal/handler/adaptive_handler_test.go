package handler

import (
	"adaptive-engine/internal/model"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockAdaptiveService
type MockAdaptiveService struct {
	mock.Mock
}

func (m *MockAdaptiveService) RecommendNextLesson(req *model.NextLessonRequest) (*model.NextLessonResponse, error) {
	args := m.Called(req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.NextLessonResponse), args.Error(1)
}

func TestNextLesson_Success(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	mockService := new(MockAdaptiveService)
	handler := NewAdaptiveHandler(mockService)
	router := gin.Default()
	router.POST("/next-lesson", handler.NextLesson)

	reqBody := &model.NextLessonRequest{
		UserID:       "user1",
		CurrentSkill: "math",
	}
	jsonBody, _ := json.Marshal(reqBody)

	expectedResp := &model.NextLessonResponse{
		NextLessonID: 123,
		Reason:       "Continue learning",
		MasteryScore: 75,
		ContentType:  "standard",
	}

	mockService.On("RecommendNextLesson", reqBody).Return(expectedResp, nil)

	// Execute
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/next-lesson", bytes.NewBuffer(jsonBody))
	router.ServeHTTP(w, req)

	// Verify
	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestNextLesson_BadRequest(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	mockService := new(MockAdaptiveService)
	handler := NewAdaptiveHandler(mockService)
	router := gin.Default()
	router.POST("/next-lesson", handler.NextLesson)

	// Invalid JSON
	reqBody := []byte(`{"invalid": "json"}`)

	// Execute
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/next-lesson", bytes.NewBuffer(reqBody))
	router.ServeHTTP(w, req)

	// Verify
	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockService.AssertNotCalled(t, "RecommendNextLesson")
}

func TestNextLesson_ServiceError(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	mockService := new(MockAdaptiveService)
	handler := NewAdaptiveHandler(mockService)
	router := gin.Default()
	router.POST("/next-lesson", handler.NextLesson)

	reqBody := &model.NextLessonRequest{
		UserID:       "user1",
		CurrentSkill: "math",
	}
	jsonBody, _ := json.Marshal(reqBody)

	mockService.On("RecommendNextLesson", reqBody).Return(nil, errors.New("service error"))

	// Execute
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/next-lesson", bytes.NewBuffer(jsonBody))
	router.ServeHTTP(w, req)

	// Verify
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockService.AssertExpectations(t)
}
