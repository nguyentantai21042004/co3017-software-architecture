package curl

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestLearnerServiceClient_GetMastery_Success(t *testing.T) {
	// Setup mock server
	expectedMastery := MasteryResponse{
		ErrorCode: 0,
		Message:   "Success",
		Data: MasteryData{
			UserID:       "user_01",
			SkillTag:     "math_algebra",
			MasteryScore: 75,
			LastUpdated:  "2025-01-01T00:00:00Z",
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/internal/learner/user_01/mastery" {
			t.Errorf("Expected path /internal/learner/user_01/mastery, got %s", r.URL.Path)
		}
		if r.URL.Query().Get("skill") != "math_algebra" {
			t.Errorf("Expected skill=math_algebra, got %s", r.URL.Query().Get("skill"))
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedMastery)
	}))
	defer server.Close()

	// Create client
	client := NewLearnerServiceClient(server.URL)

	// Test
	ctx := context.Background()
	result, err := client.GetMastery(ctx, "user_01", "math_algebra")

	// Assertions
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("Expected result, got nil")
	}
	if result.ErrorCode != expectedMastery.ErrorCode {
		t.Errorf("Expected ErrorCode %d, got %d", expectedMastery.ErrorCode, result.ErrorCode)
	}
	if result.Data.UserID != expectedMastery.Data.UserID {
		t.Errorf("Expected UserID %s, got %s", expectedMastery.Data.UserID, result.Data.UserID)
	}
	if result.Data.SkillTag != expectedMastery.Data.SkillTag {
		t.Errorf("Expected SkillTag %s, got %s", expectedMastery.Data.SkillTag, result.Data.SkillTag)
	}
	if result.Data.MasteryScore != expectedMastery.Data.MasteryScore {
		t.Errorf("Expected MasteryScore %d, got %d", expectedMastery.Data.MasteryScore, result.Data.MasteryScore)
	}
}

func TestLearnerServiceClient_GetMastery_NotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client := NewLearnerServiceClient(server.URL)
	ctx := context.Background()

	result, err := client.GetMastery(ctx, "user_01", "math_algebra")

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

func TestLearnerServiceClient_GetMastery_InvalidResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("invalid json"))
	}))
	defer server.Close()

	client := NewLearnerServiceClient(server.URL)
	ctx := context.Background()

	result, err := client.GetMastery(ctx, "user_01", "math_algebra")

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

func TestLearnerServiceClient_GetMastery_ServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := NewLearnerServiceClient(server.URL)
	ctx := context.Background()

	result, err := client.GetMastery(ctx, "user_01", "math_algebra")

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

func TestLearnerServiceClient_GetMastery_Timeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(15 * time.Second) // Longer than client timeout
	}))
	defer server.Close()

	client := NewLearnerServiceClient(server.URL)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	result, err := client.GetMastery(ctx, "user_01", "math_algebra")

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if result != nil {
		t.Errorf("Expected nil result, got %+v", result)
	}
}

// Helper functions
func IsNotFound(err error) bool {
	return err != nil && (err == ErrNotFound || contains(err.Error(), "not found"))
}

func IsInvalidResponse(err error) bool {
	return err != nil && (err == ErrInvalidResponse || contains(err.Error(), "invalid response"))
}

func IsServiceUnavailable(err error) bool {
	return err != nil && (err == ErrServiceUnavailable || contains(err.Error(), "service unavailable"))
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr ||
		(len(s) > len(substr) && (s[:len(substr)] == substr ||
			s[len(s)-len(substr):] == substr ||
			(len(s) > 1 && contains(s[1:], substr)))))
}
