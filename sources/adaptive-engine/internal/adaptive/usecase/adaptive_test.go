package usecase

import (
	"context"
	"errors"
	"testing"

	"adaptive-engine/internal/adaptive"
	"adaptive-engine/pkg/curl"
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

// Mock Learner Service Client
type mockLearnerClient struct {
	mastery *curl.MasteryResponse
	err     error
}

func (m *mockLearnerClient) GetMastery(ctx context.Context, userID, skillTag string) (*curl.MasteryResponse, error) {
	return m.mastery, m.err
}

// Mock Content Service Client
type mockContentClient struct {
	content *curl.ContentResponse
	err     error
}

func (m *mockContentClient) GetRecommendation(ctx context.Context, skillTag, contentType, userID string) (*curl.ContentResponse, error) {
	return m.content, m.err
}

func TestUsecase_RecommendNextLesson_Success_Remedial(t *testing.T) {
	ctx := context.Background()
	logger := &mockLogger{}

	masteryResp := &curl.MasteryResponse{
		ErrorCode: 0,
		Message:   "Success",
		Data: curl.MasteryData{
			UserID:       "user_01",
			SkillTag:     "math_algebra",
			MasteryScore: 30, // Below threshold
			LastUpdated:  "2025-01-01T00:00:00Z",
		},
	}

	contentResp := &curl.ContentResponse{
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
			ID:              2,
			Content:         "What is 2+2?",
			CorrectAnswer:   "4",
			SkillTag:        "math_algebra",
			DifficultyLevel: 1,
			IsRemedial:      true,
		},
	}

	learnerClient := &mockLearnerClient{mastery: masteryResp}
	contentClient := &mockContentClient{content: contentResp}

	uc := New(logger, learnerClient, contentClient)

	input := adaptive.RecommendInput{
		UserID:       "user_01",
		CurrentSkill: "math_algebra",
	}

	result, err := uc.RecommendNextLesson(ctx, input)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.NextLessonID != 2 {
		t.Errorf("Expected NextLessonID 2, got %d", result.NextLessonID)
	}
	if result.ContentType != "remedial" {
		t.Errorf("Expected ContentType 'remedial', got %s", result.ContentType)
	}
	if result.MasteryScore != 30 {
		t.Errorf("Expected MasteryScore 30, got %d", result.MasteryScore)
	}
}

func TestUsecase_RecommendNextLesson_Success_Standard(t *testing.T) {
	ctx := context.Background()
	logger := &mockLogger{}

	masteryResp := &curl.MasteryResponse{
		ErrorCode: 0,
		Message:   "Success",
		Data: curl.MasteryData{
			UserID:       "user_01",
			SkillTag:     "math_algebra",
			MasteryScore: 75, // Above threshold
			LastUpdated:  "2025-01-01T00:00:00Z",
		},
	}

	contentResp := &curl.ContentResponse{
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
			ID:              5,
			Content:         "Solve: x^2 + 5x + 6 = 0",
			CorrectAnswer:   "x = -2 or x = -3",
			SkillTag:        "math_algebra",
			DifficultyLevel: 3,
			IsRemedial:      false,
		},
	}

	learnerClient := &mockLearnerClient{mastery: masteryResp}
	contentClient := &mockContentClient{content: contentResp}

	uc := New(logger, learnerClient, contentClient)

	input := adaptive.RecommendInput{
		UserID:       "user_01",
		CurrentSkill: "math_algebra",
	}

	result, err := uc.RecommendNextLesson(ctx, input)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.NextLessonID != 5 {
		t.Errorf("Expected NextLessonID 5, got %d", result.NextLessonID)
	}
	if result.ContentType != "standard" {
		t.Errorf("Expected ContentType 'standard', got %s", result.ContentType)
	}
	if result.MasteryScore != 75 {
		t.Errorf("Expected MasteryScore 75, got %d", result.MasteryScore)
	}
}

func TestUsecase_RecommendNextLesson_FetchMasteryError(t *testing.T) {
	ctx := context.Background()
	logger := &mockLogger{}

	learnerClient := &mockLearnerClient{
		err: errors.New("service unavailable"),
	}
	contentClient := &mockContentClient{}

	uc := New(logger, learnerClient, contentClient)

	input := adaptive.RecommendInput{
		UserID:       "user_01",
		CurrentSkill: "math_algebra",
	}

	result, err := uc.RecommendNextLesson(ctx, input)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if result.NextLessonID != 0 {
		t.Errorf("Expected empty result, got %+v", result)
	}
}

func TestUsecase_RecommendNextLesson_FetchContentError(t *testing.T) {
	ctx := context.Background()
	logger := &mockLogger{}

	masteryResp := &curl.MasteryResponse{
		ErrorCode: 0,
		Message:   "Success",
		Data: curl.MasteryData{
			UserID:       "user_01",
			SkillTag:     "math_algebra",
			MasteryScore: 30,
			LastUpdated:  "2025-01-01T00:00:00Z",
		},
	}

	learnerClient := &mockLearnerClient{mastery: masteryResp}
	contentClient := &mockContentClient{
		err: errors.New("content service unavailable"),
	}

	uc := New(logger, learnerClient, contentClient)

	input := adaptive.RecommendInput{
		UserID:       "user_01",
		CurrentSkill: "math_algebra",
	}

	result, err := uc.RecommendNextLesson(ctx, input)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if result.NextLessonID != 0 {
		t.Errorf("Expected empty result, got %+v", result)
	}
}
