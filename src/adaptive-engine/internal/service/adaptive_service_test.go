package service

import (
	"adaptive-engine-service/internal/model"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecommendNextLesson_LowMastery_ReturnsRemedial(t *testing.T) {
	// Mock Learner Service
	learnerServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mastery := model.MasteryResponse{
			UserID:       "user1",
			SkillTag:     "math",
			MasteryScore: 30, // Below threshold
		}
		json.NewEncoder(w).Encode(mastery)
	}))
	defer learnerServer.Close()

	// Mock Content Service
	contentServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		question := model.QuestionResponse{
			Success: true,
			Data: model.QuestionData{
				ID:      123,
				Content: "Remedial question",
			},
		}
		json.NewEncoder(w).Encode(question)
	}))
	defer contentServer.Close()

	// Setup service
	service := NewAdaptiveService(learnerServer.URL, contentServer.URL)

	// Execute
	req := &model.NextLessonRequest{
		UserID:       "user1",
		CurrentSkill: "math",
	}
	resp, err := service.RecommendNextLesson(req)

	// Verify
	assert.NoError(t, err)
	assert.Equal(t, 123, resp.NextLessonID)
	assert.Equal(t, "remedial", resp.ContentType)
	assert.Equal(t, 30, resp.MasteryScore)
	assert.Contains(t, resp.Reason, "review the basics")
}

func TestRecommendNextLesson_HighMastery_ReturnsStandard(t *testing.T) {
	// Mock Learner Service
	learnerServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mastery := model.MasteryResponse{
			UserID:       "user1",
			SkillTag:     "math",
			MasteryScore: 75, // Above threshold
		}
		json.NewEncoder(w).Encode(mastery)
	}))
	defer learnerServer.Close()

	// Mock Content Service
	contentServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		question := model.QuestionResponse{
			Success: true,
			Data: model.QuestionData{
				ID:      456,
				Content: "Standard question",
			},
		}
		json.NewEncoder(w).Encode(question)
	}))
	defer contentServer.Close()

	// Setup service
	service := NewAdaptiveService(learnerServer.URL, contentServer.URL)

	// Execute
	req := &model.NextLessonRequest{
		UserID:       "user1",
		CurrentSkill: "math",
	}
	resp, err := service.RecommendNextLesson(req)

	// Verify
	assert.NoError(t, err)
	assert.Equal(t, 456, resp.NextLessonID)
	assert.Equal(t, "standard", resp.ContentType)
	assert.Equal(t, 75, resp.MasteryScore)
	assert.Contains(t, resp.Reason, "next challenge")
}

func TestRecommendNextLesson_ExactlyThreshold_ReturnsStandard(t *testing.T) {
	// Mock Learner Service
	learnerServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mastery := model.MasteryResponse{
			UserID:       "user1",
			SkillTag:     "math",
			MasteryScore: 50, // Exactly at threshold
		}
		json.NewEncoder(w).Encode(mastery)
	}))
	defer learnerServer.Close()

	// Mock Content Service
	contentServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		question := model.QuestionResponse{
			Success: true,
			Data: model.QuestionData{
				ID:      789,
				Content: "Standard question",
			},
		}
		json.NewEncoder(w).Encode(question)
	}))
	defer contentServer.Close()

	// Setup service
	service := NewAdaptiveService(learnerServer.URL, contentServer.URL)

	// Execute
	req := &model.NextLessonRequest{
		UserID:       "user1",
		CurrentSkill: "math",
	}
	resp, err := service.RecommendNextLesson(req)

	// Verify
	assert.NoError(t, err)
	assert.Equal(t, 789, resp.NextLessonID)
	assert.Equal(t, "standard", resp.ContentType)
	assert.Equal(t, 50, resp.MasteryScore)
}
