package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"learner-model-service/internal/learner"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUseCase is a mock implementation of learner.UseCase, defined in test_mocks.go
// MockLogger is a mock implementation of log.Logger, defined in test_mocks.go
// NewMockLogger is defined in test_mocks.go

// SetupTestRouter configures a gin router for integration testing,
// mapping only the specific API routes needed for these tests.
// This is now defined in test_mocks.go.

func TestGetMastery_Success_Integration(t *testing.T) {
	// Arrange
	mockUC := new(MockUseCase)
	mockLogger := NewMockLogger() 

	router := SetupIntegrationTestRouter(mockUC, mockLogger)
	
	userID := "user123"
	skillTag := "math_algebra"

	expectedMasteryOutput := learner.MasteryOutput{
		UserID:       userID,
		SkillTag:     skillTag,
		MasteryScore: 75,
		LastUpdated:  time.Now().Format(time.RFC3339),
	}

	// Mock expectations for the UseCase call
	mockUC.On("GetMastery", mock.Anything, learner.GetMasteryInput{UserID: userID, SkillTag: skillTag}).Return(expectedMasteryOutput, nil)

	// Act
	req := httptest.NewRequest(http.MethodGet, "/internal/learner/mastery/user123?skill=math_algebra", nil)
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

func TestGetMastery_MissingSkillTag_Integration(t *testing.T) {
	// Arrange
	mockUC := new(MockUseCase)
	mockLogger := NewMockLogger()

	router := SetupIntegrationTestRouter(mockUC, mockLogger)
	
	// Act - missing skill query param
	req := httptest.NewRequest(http.MethodGet, "/internal/learner/mastery/user123", nil)
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
	assert.Contains(t, response.Message, "invalid skill_tag: must be non-empty string") 

	mockUC.AssertNotCalled(t, "GetMastery", mock.Anything, mock.Anything)
}

func TestGetMastery_UseCaseError_Integration(t *testing.T) {
	// Arrange
	mockUC := new(MockUseCase)
	mockLogger := NewMockLogger()

	router := SetupIntegrationTestRouter(mockUC, mockLogger)
	
	expectedUseCaseError := errors.New("database error during GetMastery")
	mockUC.On("GetMastery", mock.Anything, mock.Anything).Return(learner.MasteryOutput{}, expectedUseCaseError)

	// Act
	req := httptest.NewRequest(http.MethodGet, "/internal/learner/mastery/user123?skill=math_algebra", nil)
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
	assert.Contains(t, response.Message, "failed to get mastery level: internal service error")

	mockUC.AssertExpectations(t)
}

func TestHealth_Success_Integration(t *testing.T) {
	// Arrange
	mockUC := new(MockUseCase)
	mockLogger := NewMockLogger() 

	router := SetupIntegrationTestRouter(mockUC, mockLogger)
	
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
