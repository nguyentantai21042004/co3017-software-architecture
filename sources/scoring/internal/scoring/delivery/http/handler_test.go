package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"errors" 

	"scoring/internal/scoring"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// The MockUseCase, MockLogger and SetupTestRouter are now defined in test_mocks.go

func TestSubmitAnswer_Success(t *testing.T) {
	mockLogger := NewMockLogger()
	mockUC := new(MockUseCase)

	expectedOutput := scoring.SubmitOutput{
		Correct:  true,
		Score:    100,
		Feedback: "Correct! Well done.",
	}
	mockUC.On("SubmitAnswer", mock.Anything, mock.Anything).Return(expectedOutput, nil)

	router := SetupTestRouter(mockUC, mockLogger)
	
	reqBody := SubmitRequest{
		UserID:     "user123",
		QuestionID: 1,
		Answer:     "4",
	}
	body, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", "/api/scoring/submit", bytes.NewBuffer(body))
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
	mockLogger := NewMockLogger()
	mockUC := new(MockUseCase)

	router := SetupTestRouter(mockUC, mockLogger)
	
	req, _ := http.NewRequest("POST", "/api/scoring/submit", bytes.NewBufferString("{invalid json"))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, float64(400), response["error_code"])
	assert.Contains(t, response["message"], ErrMsgBindRequestFailed) // Updated assertion
}

func TestSubmitAnswer_EmptyUserID(t *testing.T) {
	mockLogger := NewMockLogger()
	mockUC := new(MockUseCase)

	router := SetupTestRouter(mockUC, mockLogger)
	
	reqBody := SubmitRequest{
		UserID:     "",
		QuestionID: 1,
		Answer:     "4",
	}
	body, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", "/api/scoring/submit", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, float64(400), response["error_code"])
	assert.Contains(t, response["message"], ErrMsgInvalidUserID) // Updated assertion
}

func TestSubmitAnswer_InvalidQuestionID(t *testing.T) {
	mockLogger := NewMockLogger()
	mockUC := new(MockUseCase)

	router := SetupTestRouter(mockUC, mockLogger)
	
	reqBody := SubmitRequest{
		UserID:     "user123",
		QuestionID: 0,
		Answer:     "4",
	}
	body, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", "/api/scoring/submit", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, float64(400), response["error_code"])
	assert.Contains(t, response["message"], ErrMsgInvalidQuestionID) // Updated assertion
}

func TestSubmitAnswer_EmptyAnswer(t *testing.T) {
	mockLogger := NewMockLogger()
	mockUC := new(MockUseCase)

	router := SetupTestRouter(mockUC, mockLogger)
	
	reqBody := SubmitRequest{
		UserID:     "user123",
		QuestionID: 1,
		Answer:     "",
	}
	body, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", "/api/scoring/submit", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, float64(400), response["error_code"])
	assert.Contains(t, response["message"], ErrMsgInvalidAnswer) // Updated assertion
}

func TestSubmitAnswer_UseCaseError(t *testing.T) {
	mockLogger := NewMockLogger()
	mockUC := new(MockUseCase)

	expectedErr := errors.New("usecase error")
	mockUC.On("SubmitAnswer", mock.Anything, mock.Anything).Return(scoring.SubmitOutput{}, expectedErr)

	router := SetupTestRouter(mockUC, mockLogger)
	
	reqBody := SubmitRequest{
		UserID:     "user123",
		QuestionID: 1,
		Answer:     "4",
	}
	body, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", "/api/scoring/submit", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, float64(400), response["error_code"])
	assert.Contains(t, response["message"], ErrMsgSubmitAnswerFailed) // Updated assertion

	mockUC.AssertExpectations(t)
}

func TestHealth(t *testing.T) {
	mockLogger := NewMockLogger()
	mockUC := new(MockUseCase)

	router := SetupTestRouter(mockUC, mockLogger)

	// Corrected URL: main.go registers /health globally, not under /api/scoring
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
	mockLogger := NewMockLogger()
	mockUC := new(MockUseCase)

	handler := New(mockLogger, mockUC) // This `handler` is used for the return value assertion

	assert.NotNil(t, handler)
}
