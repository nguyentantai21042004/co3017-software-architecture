package curl

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestContentServiceClient_GetQuestion_Success(t *testing.T) {
	// Mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/1", r.URL.Path)
		assert.Equal(t, "GET", r.Method)

		resp := ContentQuestionResponse{
			ErrorCode: 0,
			Message:   "Success",
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewContentServiceClient(server.URL)
	ctx := context.Background()

	result, err := client.GetQuestion(ctx, 1)

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 0, result.ErrorCode)
	assert.Equal(t, "Success", result.Message)
	assert.Equal(t, int64(1), result.Data.ID)
	assert.Equal(t, "4", result.Data.CorrectAnswer)
	assert.Equal(t, "math", result.Data.SkillTag)
}

func TestContentServiceClient_GetQuestion_NotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "not found"}`))
	}))
	defer server.Close()

	client := NewContentServiceClient(server.URL)
	ctx := context.Background()

	result, err := client.GetQuestion(ctx, 999)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), ErrMsgResourceNotFound)
}

func TestContentServiceClient_GetQuestion_InvalidResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{invalid json`))
	}))
	defer server.Close()

	client := NewContentServiceClient(server.URL)
	ctx := context.Background()

	result, err := client.GetQuestion(ctx, 1)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), ErrMsgResponseParseFailed)
}

func TestContentServiceClient_GetQuestion_ServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "internal server error"}`))
	}))
	defer server.Close()

	client := NewContentServiceClient(server.URL)
	ctx := context.Background()

	result, err := client.GetQuestion(ctx, 1)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), ErrMsgUnexpectedStatusCode)
}

func TestContentServiceClient_GetQuestion_Timeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second) // Longer than client timeout
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Create client with short timeout
	baseClient := NewClient(server.URL, 1*time.Second)
	client := &ContentServiceClient{client: baseClient}
	ctx := context.Background()

	result, err := client.GetQuestion(ctx, 1)

	require.Error(t, err)
	assert.Nil(t, result)
	// Timeout errors can vary, just check that error occurred
}

func TestContentServiceClient_GetQuestion_InvalidQuestionID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer server.Close()

	client := NewContentServiceClient(server.URL)
	ctx := context.Background()

	result, err := client.GetQuestion(ctx, -1)

	require.Error(t, err)
	assert.Nil(t, result)
}

func TestNewContentServiceClient(t *testing.T) {
	baseURL := "http://localhost:8081/api/content"
	client := NewContentServiceClient(baseURL)

	assert.NotNil(t, client)
	assert.NotNil(t, client.client)
	assert.Equal(t, baseURL, client.client.GetBaseURL())
}
