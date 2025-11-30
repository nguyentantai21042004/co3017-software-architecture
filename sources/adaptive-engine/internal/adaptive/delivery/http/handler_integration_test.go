package http

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"adaptive-engine/internal/adaptive"
	"adaptive-engine/pkg/curl"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUseCase for integration testing
type MockUseCase struct {
	mock.Mock
}

func (m *MockUseCase) RecommendNextLesson(ctx context.Context, input adaptive.RecommendInput) (adaptive.RecommendOutput, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		return adaptive.RecommendOutput{}, args.Error(1)
	}
	return args.Get(0).(adaptive.RecommendOutput), args.Error(1)
}

// Ensure MockUseCase implements adaptive.UseCase
var _ adaptive.UseCase = (*MockUseCase)(nil)

// MockLogger for integration testing
type MockLogger struct {
	mock.Mock
}

func (m *MockLogger) Debug(ctx context.Context, args ...any)                   {}
func (m *MockLogger) Debugf(ctx context.Context, template string, args ...any) {}
func (m *MockLogger) Info(ctx context.Context, args ...any)                    {}
func (m *MockLogger) Infof(ctx context.Context, template string, args ...any)  {}
func (m *MockLogger) Warn(ctx context.Context, args ...any)                    {}
func (m *MockLogger) Warnf(ctx context.Context, template string, args ...any)  {}
func (m *MockLogger) Error(ctx context.Context, args ...any)                   {}
func (m *MockLogger) Errorf(ctx context.Context, template string, args ...any) {}
func (m *MockLogger) Fatal(ctx context.Context, args ...any)                   {}
func (m *MockLogger) Fatalf(ctx context.Context, template string, args ...any) {}

// SetupTestRouter creates a test router with the handler
func SetupTestRouter(handler Handler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	api := router.Group("/api/adaptive")
	MapAdaptiveRoutes(api, handler)
	router.GET("/health", func(c *gin.Context) {
		handler.Health(c)
	})
	return router
}

func TestNextLesson_Integration_Success_Remedial(t *testing.T) {
	// Arrange
	mockUC := new(MockUseCase)
	mockLogger := new(MockLogger)

	expectedOutput := adaptive.RecommendOutput{
		NextLessonID: 2,
		Reason:       "Your mastery is 30%. Let's review the basics.",
		MasteryScore: 30,
		ContentType:  "remedial",
	}

	input := adaptive.RecommendInput{
		UserID:       "user_01",
		CurrentSkill: "math_algebra",
	}

	mockUC.On("RecommendNextLesson", mock.Anything, input).Return(expectedOutput, nil)

	handler := New(mockLogger, mockUC)
	router := SetupTestRouter(handler)

	reqBody := NextLessonRequest{
		UserID:       "user_01",
		CurrentSkill: "math_algebra",
	}
	bodyBytes, _ := json.Marshal(reqBody)

	// Act
	req := httptest.NewRequest("POST", "/api/adaptive/next-lesson", bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var resp map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, float64(0), resp["error_code"].(float64))

	data := resp["data"].(map[string]interface{})
	assert.Equal(t, float64(2), data["next_lesson_id"].(float64))
	assert.Equal(t, "remedial", data["content_type"].(string))
	assert.Equal(t, float64(30), data["mastery_score"].(float64))
	assert.Contains(t, data["reason"].(string), "30%")

	mockUC.AssertExpectations(t)
}

func TestNextLesson_Integration_Success_Standard(t *testing.T) {
	// Arrange
	mockUC := new(MockUseCase)
	mockLogger := new(MockLogger)

	expectedOutput := adaptive.RecommendOutput{
		NextLessonID: 5,
		Reason:       "Great! Your mastery is 75%. Continue with the next challenge.",
		MasteryScore: 75,
		ContentType:  "standard",
	}

	input := adaptive.RecommendInput{
		UserID:       "user_01",
		CurrentSkill: "math_algebra",
	}

	mockUC.On("RecommendNextLesson", mock.Anything, input).Return(expectedOutput, nil)

	handler := New(mockLogger, mockUC)
	router := SetupTestRouter(handler)

	reqBody := NextLessonRequest{
		UserID:       "user_01",
		CurrentSkill: "math_algebra",
	}
	bodyBytes, _ := json.Marshal(reqBody)

	// Act
	req := httptest.NewRequest("POST", "/api/adaptive/next-lesson", bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var resp map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, float64(0), resp["error_code"].(float64))

	data := resp["data"].(map[string]interface{})
	assert.Equal(t, float64(5), data["next_lesson_id"].(float64))
	assert.Equal(t, "standard", data["content_type"].(string))
	assert.Equal(t, float64(75), data["mastery_score"].(float64))
	assert.Contains(t, data["reason"].(string), "75%")

	mockUC.AssertExpectations(t)
}

func TestNextLesson_Integration_UseCaseError(t *testing.T) {
	// Arrange
	mockUC := new(MockUseCase)
	mockLogger := new(MockLogger)

	input := adaptive.RecommendInput{
		UserID:       "user_01",
		CurrentSkill: "math_algebra",
	}

	mockUC.On("RecommendNextLesson", mock.Anything, input).Return(adaptive.RecommendOutput{}, curl.ErrServiceUnavailable)

	handler := New(mockLogger, mockUC)
	router := SetupTestRouter(handler)

	reqBody := NextLessonRequest{
		UserID:       "user_01",
		CurrentSkill: "math_algebra",
	}
	bodyBytes, _ := json.Marshal(reqBody)

	// Act
	req := httptest.NewRequest("POST", "/api/adaptive/next-lesson", bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)

	var resp map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.NotEqual(t, float64(0), resp["error_code"].(float64))

	mockUC.AssertExpectations(t)
}

func TestNextLesson_Integration_EdgeCase_ThresholdBoundary(t *testing.T) {
	// Test case where mastery score is exactly at threshold (50)
	mockUC := new(MockUseCase)
	mockLogger := new(MockLogger)

	expectedOutput := adaptive.RecommendOutput{
		NextLessonID: 3,
		Reason:       "Great! Your mastery is 50%. Continue with the next challenge.",
		MasteryScore: 50,
		ContentType:  "standard", // Should be standard when score >= threshold
	}

	input := adaptive.RecommendInput{
		UserID:       "user_01",
		CurrentSkill: "math_algebra",
	}

	mockUC.On("RecommendNextLesson", mock.Anything, input).Return(expectedOutput, nil)

	handler := New(mockLogger, mockUC)
	router := SetupTestRouter(handler)

	reqBody := NextLessonRequest{
		UserID:       "user_01",
		CurrentSkill: "math_algebra",
	}
	bodyBytes, _ := json.Marshal(reqBody)

	// Act
	req := httptest.NewRequest("POST", "/api/adaptive/next-lesson", bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var resp map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, float64(0), resp["error_code"].(float64))

	data := resp["data"].(map[string]interface{})
	assert.Equal(t, "standard", data["content_type"].(string))
	assert.Equal(t, float64(50), data["mastery_score"].(float64))

	mockUC.AssertExpectations(t)
}

