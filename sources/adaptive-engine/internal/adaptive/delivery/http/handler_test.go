package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"adaptive-engine/internal/adaptive"

	"github.com/gin-gonic/gin"
)

// Mock logger
type mockLogger struct{}

func (m *mockLogger) Debug(ctx context.Context, args ...any)                   {}
func (m *mockLogger) Debugf(ctx context.Context, template string, args ...any) {}
func (m *mockLogger) Info(ctx context.Context, args ...any)                    {}
func (m *mockLogger) Infof(ctx context.Context, template string, args ...any)  {}
func (m *mockLogger) Warn(ctx context.Context, args ...any)                    {}
func (m *mockLogger) Warnf(ctx context.Context, template string, args ...any)  {}
func (m *mockLogger) Error(ctx context.Context, args ...any)                   {}
func (m *mockLogger) Errorf(ctx context.Context, template string, args ...any) {}
func (m *mockLogger) Fatal(ctx context.Context, args ...any)                   {}
func (m *mockLogger) Fatalf(ctx context.Context, template string, args ...any) {}

// Mock usecase
type mockUsecase struct {
	output adaptive.RecommendOutput
	err    error
}

func (m *mockUsecase) RecommendNextLesson(ctx context.Context, input adaptive.RecommendInput) (adaptive.RecommendOutput, error) {
	return m.output, m.err
}

func TestHandler_NextLesson_Success(t *testing.T) {
	logger := &mockLogger{}
	uc := &mockUsecase{
		output: adaptive.RecommendOutput{
			NextLessonID: 2,
			Reason:       "Your mastery is 30%. Let's review the basics.",
			MasteryScore: 30,
			ContentType:  "remedial",
		},
	}

	handler := New(logger, uc)

	reqBody := NextLessonRequest{
		UserID:       "user_01",
		CurrentSkill: "math_algebra",
	}
	bodyBytes, _ := json.Marshal(reqBody)

	req := httptest.NewRequest("POST", "/api/adaptive/next-lesson", bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	ginCtx, _ := gin.CreateTestContext(w)
	ginCtx.Request = req

	handler.NextLesson(ginCtx)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var resp map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	if resp["error_code"].(float64) != 0 {
		t.Errorf("Expected error_code 0, got %v", resp["error_code"])
	}

	data := resp["data"].(map[string]interface{})
	if int(data["next_lesson_id"].(float64)) != 2 {
		t.Errorf("Expected next_lesson_id 2, got %v", data["next_lesson_id"])
	}
}

func TestHandler_NextLesson_InvalidJSON(t *testing.T) {
	logger := &mockLogger{}
	uc := &mockUsecase{}

	handler := New(logger, uc)

	req := httptest.NewRequest("POST", "/api/adaptive/next-lesson", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	ginCtx, _ := gin.CreateTestContext(w)
	ginCtx.Request = req

	handler.NextLesson(ginCtx)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}

func TestHandler_NextLesson_EmptyUserID(t *testing.T) {
	logger := &mockLogger{}
	uc := &mockUsecase{}

	handler := New(logger, uc)

	reqBody := NextLessonRequest{
		UserID:       "",
		CurrentSkill: "math_algebra",
	}
	bodyBytes, _ := json.Marshal(reqBody)

	req := httptest.NewRequest("POST", "/api/adaptive/next-lesson", bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	ginCtx, _ := gin.CreateTestContext(w)
	ginCtx.Request = req

	handler.NextLesson(ginCtx)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}

func TestHandler_NextLesson_EmptySkill(t *testing.T) {
	logger := &mockLogger{}
	uc := &mockUsecase{}

	handler := New(logger, uc)

	reqBody := NextLessonRequest{
		UserID:       "user_01",
		CurrentSkill: "",
	}
	bodyBytes, _ := json.Marshal(reqBody)

	req := httptest.NewRequest("POST", "/api/adaptive/next-lesson", bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	ginCtx, _ := gin.CreateTestContext(w)
	ginCtx.Request = req

	handler.NextLesson(ginCtx)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}

func TestHandler_NextLesson_UsecaseError(t *testing.T) {
	logger := &mockLogger{}
	uc := &mockUsecase{
		err: errors.New("service unavailable"),
	}

	handler := New(logger, uc)

	reqBody := NextLessonRequest{
		UserID:       "user_01",
		CurrentSkill: "math_algebra",
	}
	bodyBytes, _ := json.Marshal(reqBody)

	req := httptest.NewRequest("POST", "/api/adaptive/next-lesson", bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	ginCtx, _ := gin.CreateTestContext(w)
	ginCtx.Request = req

	handler.NextLesson(ginCtx)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}

func TestHandler_Health(t *testing.T) {
	logger := &mockLogger{}
	uc := &mockUsecase{}

	handler := New(logger, uc)

	req := httptest.NewRequest("GET", "/api/adaptive/health", nil)
	w := httptest.NewRecorder()

	ginCtx, _ := gin.CreateTestContext(w)
	ginCtx.Request = req

	handler.Health(ginCtx)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var resp map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	if resp["error_code"].(float64) != 0 {
		t.Errorf("Expected error_code 0, got %v", resp["error_code"])
	}
}
