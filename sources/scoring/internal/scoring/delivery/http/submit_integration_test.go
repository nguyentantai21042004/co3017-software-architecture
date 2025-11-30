package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"scoring/internal/scoring"
	
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSubmitAnswer_Success_Correct(t *testing.T) {
	mockUc := new(MockUseCase)
	mockLogger := NewMockLogger() // Use the factory function

	expectedOutput := scoring.SubmitOutput{
		Correct:  true,
		Score:    100,
		Feedback: "Correct! Well done.",
	}
	mockUc.On("SubmitAnswer", mock.Anything, mock.Anything).Return(expectedOutput, nil)

	router := SetupTestRouter(mockUc, mockLogger)

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

	var respBody map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &respBody)

	assert.Equal(t, float64(0), respBody["error_code"])
	assert.Equal(t, "Success", respBody["message"])
	data := respBody["data"].(map[string]interface{})
	assert.Equal(t, true, data["correct"])
	assert.Equal(t, float64(100), data["score"])
	assert.Equal(t, "Correct! Well done.", data["feedback"])

	mockUc.AssertExpectations(t)
}

func TestSubmitAnswer_Success_Incorrect(t *testing.T) {
	mockUc := new(MockUseCase)
	mockLogger := NewMockLogger()

	expectedOutput := scoring.SubmitOutput{
		Correct:  false,
		Score:    0,
		Feedback: "Incorrect answer. Please try again!",
	}
	mockUc.On("SubmitAnswer", mock.Anything, mock.Anything).Return(expectedOutput, nil)

	router := SetupTestRouter(mockUc, mockLogger)

	reqBody := SubmitRequest{
		UserID:     "user123",
		QuestionID: 1,
		Answer:     "5",
	}
	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/api/scoring/submit", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var respBody map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &respBody)

	assert.Equal(t, float64(0), respBody["error_code"])
	assert.Equal(t, "Success", respBody["message"])
	data := respBody["data"].(map[string]interface{})
	assert.Equal(t, false, data["correct"])
	assert.Equal(t, float64(0), data["score"])
	assert.Equal(t, "Incorrect answer. Please try again!", data["feedback"])

	mockUc.AssertExpectations(t)
}

func TestSubmitAnswer_InvalidJson(t *testing.T) {
	mockUc := new(MockUseCase)
	mockLogger := NewMockLogger()

	router := SetupTestRouter(mockUc, mockLogger)

	req, _ := http.NewRequest("POST", "/api/scoring/submit", bytes.NewBuffer([]byte("{invalid json}")))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var respBody map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &respBody)

	assert.Equal(t, float64(http.StatusBadRequest), respBody["error_code"])
	assert.Equal(t, "failed to bind request body: invalid JSON format or missing required fields", respBody["message"]) // Updated assertion

	mockUc.AssertNotCalled(t, "SubmitAnswer", mock.Anything, mock.Anything)
}

func TestSubmitAnswer_MissingFields(t *testing.T) {
	mockUc := new(MockUseCase)
	mockLogger := NewMockLogger()

	router := SetupTestRouter(mockUc, mockLogger)

	// Missing UserID
	reqBody := SubmitRequest{
		QuestionID: 1,
		Answer:     "4",
	}
	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/api/scoring/submit", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	var respBody map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &respBody)
	assert.Equal(t, float64(http.StatusBadRequest), respBody["error_code"])
	assert.Equal(t, "invalid user_id: must be non-empty string", respBody["message"]) // Updated assertion

	// Missing QuestionID
	reqBody = SubmitRequest{
		UserID: "user123",
		Answer: "4",
	}
	body, _ = json.Marshal(reqBody)
	req, _ = http.NewRequest("POST", "/api/scoring/submit", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	json.Unmarshal(w.Body.Bytes(), &respBody)
	assert.Equal(t, float64(http.StatusBadRequest), respBody["error_code"])
	assert.Equal(t, "invalid question_id: must be positive integer", respBody["message"]) // Updated assertion

	// Missing Answer
	reqBody = SubmitRequest{
		UserID:     "user123",
		QuestionID: 1,
	}
	body, _ = json.Marshal(reqBody)
	req, _ = http.NewRequest("POST", "/api/scoring/submit", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	json.Unmarshal(w.Body.Bytes(), &respBody)
	assert.Equal(t, float64(http.StatusBadRequest), respBody["error_code"])
	assert.Equal(t, "invalid answer: must be non-empty string", respBody["message"]) // Updated assertion

	mockUc.AssertNotCalled(t, "SubmitAnswer", mock.Anything, mock.Anything)
}

func TestSubmitAnswer_UseCaseErrorInIntegration(t *testing.T) { // Renamed to avoid collision
	mockUc := new(MockUseCase)
	mockLogger := NewMockLogger()

	expectedError := errors.New("something went wrong in use case")
	mockUc.On("SubmitAnswer", mock.Anything, mock.Anything).Return(scoring.SubmitOutput{}, expectedError)

	router := SetupTestRouter(mockUc, mockLogger)

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

	assert.Equal(t, http.StatusBadRequest, w.Code) // The handler returns 400 for use case errors
	var respBody map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &respBody)
	assert.Equal(t, float64(http.StatusBadRequest), respBody["error_code"])
	assert.Equal(t, "failed to submit answer: internal service error", respBody["message"]) // Updated assertion

	mockUc.AssertExpectations(t)
}

// Test case for GetAnsweredQuestions (called by content-service, not directly an endpoint of scoring)
func TestGetAnsweredQuestions_Success(t *testing.T) {
	mockUc := new(MockUseCase)
	mockLogger := NewMockLogger()

	expectedQuestionIDs := []int64{101, 102, 103}
	mockUc.On("GetAnsweredQuestions", mock.Anything, "user456", "math").Return(expectedQuestionIDs, nil)

	// Since this is called by content-service, the endpoint path is /api/scoring/answered-questions
	router := SetupTestRouter(mockUc, mockLogger)
	
	req, _ := http.NewRequest("GET", "/api/scoring/answered-questions?user_id=user456&skill=math", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var respBody map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &respBody)

	assert.Equal(t, float64(0), respBody["error_code"])
	assert.Equal(t, "Success", respBody["message"])
	data := respBody["data"].([]interface{}) // JSON unmarshals numbers as float64
	assert.Len(t, data, 3)
	assert.Equal(t, float64(101), data[0])
	assert.Equal(t, float64(102), data[1])
	assert.Equal(t, float64(103), data[2])

	mockUc.AssertExpectations(t)
}

func TestGetAnsweredQuestions_UseCaseError(t *testing.T) {
	mockUc := new(MockUseCase)
	mockLogger := NewMockLogger()

	expectedError := errors.New("failed to get answered questions from use case")
	mockUc.On("GetAnsweredQuestions", mock.Anything, "user456", "math").Return(nil, expectedError)

	router := SetupTestRouter(mockUc, mockLogger)
	
	req, _ := http.NewRequest("GET", "/api/scoring/answered-questions?user_id=user456&skill=math", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code) // Assuming generic error handling returns 500

	var respBody map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &respBody)

	assert.Equal(t, float64(http.StatusInternalServerError), respBody["error_code"])
	assert.Equal(t, "failed to retrieve answered questions", respBody["message"]) // Updated assertion

	mockUc.AssertExpectations(t)
}
