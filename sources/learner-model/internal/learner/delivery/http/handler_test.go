package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"learner-model-service/internal/learner"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUseCase is a mock implementation of learner.UseCase
type MockUseCase struct {
	mock.Mock
}

func (m *MockUseCase) GetMastery(ctx context.Context, input learner.GetMasteryInput) (learner.MasteryOutput, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(learner.MasteryOutput), args.Error(1)
}

func (m *MockUseCase) UpdateMasteryFromEvent(ctx context.Context, input learner.UpdateMasteryInput) error {
	args := m.Called(ctx, input)
	return args.Error(0)
}

// MockLogger for testing
type MockLogger struct {
	mock.Mock
}

func (m *MockLogger) Debug(ctx context.Context, arg ...any) { m.Called(ctx, arg) }
func (m *MockLogger) Debugf(ctx context.Context, format string, arg ...any) {
	m.Called(ctx, format, arg)
}
func (m *MockLogger) Info(ctx context.Context, arg ...any) { m.Called(ctx, arg) }
func (m *MockLogger) Infof(ctx context.Context, format string, arg ...any) {
	m.Called(ctx, format, arg)
}
func (m *MockLogger) Warn(ctx context.Context, arg ...any) { m.Called(ctx, arg) }
func (m *MockLogger) Warnf(ctx context.Context, format string, arg ...any) {
	m.Called(ctx, format, arg)
}
func (m *MockLogger) Error(ctx context.Context, arg ...any) { m.Called(ctx, arg) }
func (m *MockLogger) Errorf(ctx context.Context, format string, arg ...any) {
	m.Called(ctx, format, arg)
}
func (m *MockLogger) Fatal(ctx context.Context, arg ...any) { m.Called(ctx, arg) }
func (m *MockLogger) Fatalf(ctx context.Context, format string, arg ...any) {
	m.Called(ctx, format, arg)
}

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.New()
}

func TestGetMastery_Success(t *testing.T) {
	// Arrange
	mockUC := new(MockUseCase)
	mockLogger := new(MockLogger)

	h := &handler{
		l:  mockLogger,
		uc: mockUC,
	}

	router := setupTestRouter()
	router.GET("/mastery/:user_id", func(c *gin.Context) {
		h.GetMastery(c)
	})

	userID := "user123"
	skillTag := "math_algebra"

	expectedOutput := learner.MasteryOutput{
		UserID:       userID,
		SkillTag:     skillTag,
		MasteryScore: 75,
		LastUpdated:  "2024-01-01T00:00:00Z",
	}

	// Mock expectations
	mockUC.On("GetMastery", mock.Anything, mock.MatchedBy(func(input learner.GetMasteryInput) bool {
		return input.UserID == userID && input.SkillTag == skillTag
	})).Return(expectedOutput, nil)

	// Act
	req := httptest.NewRequest(http.MethodGet, "/mastery/user123?skill=math_algebra", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var response struct {
		ErrorCode int             `json:"error_code"`
		Message   string          `json:"message"`
		Data      MasteryResponse `json:"data"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, 0, response.ErrorCode)
	assert.Equal(t, "Success", response.Message)
	assert.Equal(t, userID, response.Data.UserID)
	assert.Equal(t, skillTag, response.Data.SkillTag)
	assert.Equal(t, 75, response.Data.MasteryScore)

	mockUC.AssertExpectations(t)
}

// Note: Testing missing user_id is difficult with Gin's path params
// The route won't match if user_id is truly missing, resulting in 404
// We test missing skillTag instead which is a query param

func TestGetMastery_MissingSkillTag(t *testing.T) {
	// Arrange
	mockUC := new(MockUseCase)
	mockLogger := new(MockLogger)

	// Mock logger expectations
	mockLogger.On("Errorf", mock.Anything, mock.Anything, mock.Anything).Return()

	h := &handler{
		l:  mockLogger,
		uc: mockUC,
	}

	router := setupTestRouter()
	router.GET("/mastery/:user_id", func(c *gin.Context) {
		h.GetMastery(c)
	})

	// Act - missing skill query param
	req := httptest.NewRequest(http.MethodGet, "/mastery/user123", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response struct {
		ErrorCode int    `json:"error_code"`
		Message   string `json:"message"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, 400, response.ErrorCode)
	assert.Contains(t, response.Message, "skill_tag")

	mockUC.AssertNotCalled(t, "GetMastery", mock.Anything, mock.Anything)
}

func TestGetMastery_UseCaseError(t *testing.T) {
	// Arrange
	mockUC := new(MockUseCase)
	mockLogger := new(MockLogger)

	h := &handler{
		l:  mockLogger,
		uc: mockUC,
	}

	router := setupTestRouter()
	router.GET("/mastery/:user_id", func(c *gin.Context) {
		h.GetMastery(c)
	})

	// Mock error
	mockLogger.On("Errorf", mock.Anything, mock.Anything, mock.Anything).Return()
	mockUC.On("GetMastery", mock.Anything, mock.Anything).Return(learner.MasteryOutput{}, errors.New("database error"))

	// Act
	req := httptest.NewRequest(http.MethodGet, "/mastery/user123?skill=math_algebra", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var response struct {
		ErrorCode int    `json:"error_code"`
		Message   string `json:"message"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, 500, response.ErrorCode)
	assert.Contains(t, response.Message, "failed to get mastery")

	mockUC.AssertExpectations(t)
}

func TestHealth_Success(t *testing.T) {
	// Arrange
	mockUC := new(MockUseCase)
	mockLogger := new(MockLogger)

	h := &handler{
		l:  mockLogger,
		uc: mockUC,
	}

	router := setupTestRouter()
	router.GET("/health", func(c *gin.Context) {
		h.Health(c)
	})

	// Act
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var response struct {
		ErrorCode int                    `json:"error_code"`
		Message   string                 `json:"message"`
		Data      map[string]interface{} `json:"data"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, 0, response.ErrorCode)
	assert.Equal(t, "Healthy", response.Message)
	assert.Equal(t, "healthy", response.Data["status"])
	assert.Equal(t, "learner-model", response.Data["service"])
}

func TestToMasteryResponse(t *testing.T) {
	// Arrange
	output := learner.MasteryOutput{
		UserID:       "user456",
		SkillTag:     "science_physics",
		MasteryScore: 85,
		LastUpdated:  "2024-01-15T10:30:00Z",
	}

	// Act
	response := ToMasteryResponse(output)

	// Assert
	assert.Equal(t, "user456", response.UserID)
	assert.Equal(t, "science_physics", response.SkillTag)
	assert.Equal(t, 85, response.MasteryScore)
	assert.Equal(t, "2024-01-15T10:30:00Z", response.LastUpdated)
}

func TestGetMasteryRequest_ToGetMasteryInput(t *testing.T) {
	// Arrange
	req := GetMasteryRequest{
		UserID:   "user789",
		SkillTag: "english_grammar",
	}

	// Act
	input := req.ToGetMasteryInput()

	// Assert
	assert.Equal(t, "user789", input.UserID)
	assert.Equal(t, "english_grammar", input.SkillTag)
}
