package curl

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestContentServiceClient_GetRecommendation_Success(t *testing.T) {
	expectedContent := ContentResponse{
		Success: true,
		Message: "Content found",
		Data: struct {
			ID              int64  `json:"id"`
			Content         string `json:"content"`
			CorrectAnswer   string `json:"correct_answer"`
			SkillTag        string `json:"skill_tag"`
			DifficultyLevel int    `json:"difficulty_level"`
			IsRemedial      bool   `json:"is_remedial"`
		}{
			ID:              1,
			Content:         "What is 2+2?",
			CorrectAnswer:   "4",
			SkillTag:        "math_algebra",
			DifficultyLevel: 1,
			IsRemedial:      true,
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/api/content/recommend" {
			t.Errorf("Expected path /api/content/recommend, got %s", r.URL.Path)
		}
		if r.URL.Query().Get("skill") != "math_algebra" {
			t.Errorf("Expected skill=math_algebra, got %s", r.URL.Query().Get("skill"))
		}
		if r.URL.Query().Get("type") != "remedial" {
			t.Errorf("Expected type=remedial, got %s", r.URL.Query().Get("type"))
		}
		if r.URL.Query().Get("userId") != "user_01" {
			t.Errorf("Expected userId=user_01, got %s", r.URL.Query().Get("userId"))
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedContent)
	}))
	defer server.Close()

	client := NewContentServiceClient(server.URL)
	ctx := context.Background()

	result, err := client.GetRecommendation(ctx, "math_algebra", "remedial", "user_01")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("Expected result, got nil")
	}
	if result.Data.ID != expectedContent.Data.ID {
		t.Errorf("Expected ID %d, got %d", expectedContent.Data.ID, result.Data.ID)
	}
	if result.Data.SkillTag != expectedContent.Data.SkillTag {
		t.Errorf("Expected SkillTag %s, got %s", expectedContent.Data.SkillTag, result.Data.SkillTag)
	}
}

func TestContentServiceClient_GetRecommendation_NotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client := NewContentServiceClient(server.URL)
	ctx := context.Background()

	result, err := client.GetRecommendation(ctx, "math_algebra", "remedial", "user_01")

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if result != nil {
		t.Errorf("Expected nil result, got %+v", result)
	}
	if !IsNotFound(err) {
		t.Errorf("Expected NotFound error, got %v", err)
	}
}

func TestContentServiceClient_GetRecommendation_InvalidResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("invalid json"))
	}))
	defer server.Close()

	client := NewContentServiceClient(server.URL)
	ctx := context.Background()

	result, err := client.GetRecommendation(ctx, "math_algebra", "remedial", "user_01")

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if result != nil {
		t.Errorf("Expected nil result, got %+v", result)
	}
	if !IsInvalidResponse(err) {
		t.Errorf("Expected InvalidResponse error, got %v", err)
	}
}

func TestContentServiceClient_GetRecommendation_ServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := NewContentServiceClient(server.URL)
	ctx := context.Background()

	result, err := client.GetRecommendation(ctx, "math_algebra", "remedial", "user_01")

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if result != nil {
		t.Errorf("Expected nil result, got %+v", result)
	}
	if !IsServiceUnavailable(err) {
		t.Errorf("Expected ServiceUnavailable error, got %v", err)
	}
}
