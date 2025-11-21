package system

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const AdaptiveEngineURL = "http://localhost:8084"

type NextLessonRequest struct {
	UserID       string `json:"user_id"`
	CurrentSkill string `json:"current_skill"`
}

type NextLessonResponse struct {
	NextLessonID int    `json:"next_lesson_id"`
	ContentType  string `json:"content_type"`
	Reason       string `json:"reason"`
	MasteryScore int    `json:"mastery_score"`
}

// TC_ADAPTIVE_01: Trigger Remediation (Weak Student)
func TestAdaptiveRemediation(t *testing.T) {
	fmt.Println("\n🧪 TC_ADAPTIVE_01: Trigger Remediation (Weak Student)")

	// Setup: Create a user with low mastery by submitting wrong answers
	userID := "test-adaptive-weak"

	// Submit 2 wrong answers to get low mastery
	for i := 0; i < 2; i++ {
		req := SubmitAnswerRequest{
			UserID:     userID,
			QuestionID: 6,
			Answer:     "WRONG",
		}
		submitAnswer(req)
	}

	time.Sleep(2 * time.Second)

	// Get next lesson recommendation
	nextLesson, err := getNextLesson(userID, "math")
	require.NoError(t, err, "Failed to get next lesson")

	// Verify remedial content is recommended
	assert.Equal(t, "remedial", nextLesson.ContentType, "Should recommend remedial for weak student")
	assert.Less(t, nextLesson.MasteryScore, 50, "Mastery should be below threshold")

	fmt.Printf("✅ Remediation triggered: mastery=%d%%, type=%s\n",
		nextLesson.MasteryScore, nextLesson.ContentType)
}

// TC_ADAPTIVE_02: Trigger Advancement (Strong Student)
func TestAdaptiveAdvancement(t *testing.T) {
	fmt.Println("\n🧪 TC_ADAPTIVE_02: Trigger Advancement (Strong Student)")

	// Setup: Create a user with high mastery by submitting correct answers
	userID := "test-adaptive-strong"

	// Get question first
	question, _ := getQuestion(6)

	// Submit 3 correct answers to get high mastery
	for i := 0; i < 3; i++ {
		req := SubmitAnswerRequest{
			UserID:     userID,
			QuestionID: 6,
			Answer:     question.CorrectAnswer,
		}
		submitAnswer(req)
	}

	time.Sleep(2 * time.Second)

	// Get next lesson recommendation
	nextLesson, err := getNextLesson(userID, "math")
	require.NoError(t, err, "Failed to get next lesson")

	// Verify standard content is recommended
	assert.Equal(t, "standard", nextLesson.ContentType, "Should recommend standard for strong student")
	assert.GreaterOrEqual(t, nextLesson.MasteryScore, 50, "Mastery should be at/above threshold")

	fmt.Printf("✅ Advancement triggered: mastery=%d%%, type=%s\n",
		nextLesson.MasteryScore, nextLesson.ContentType)
}

// TC_ADAPTIVE_03: Learner Service Unavailable (Fallback)
func TestAdaptiveFallback(t *testing.T) {
	fmt.Println("\n🧪 TC_ADAPTIVE_03: Learner Service Unavailable (Fallback)")
	fmt.Println("⚠️  This test requires manual Learner Model service stop")
	fmt.Println("⚠️  Skipping automated execution - verify manually:")
	fmt.Println("   1. Stop Learner Model service")
	fmt.Println("   2. Call /adaptive/next-lesson")
	fmt.Println("   3. Should return 200 with default/remedial content")
	fmt.Println("   4. Should NOT crash with 500")

	t.Skip("Manual test - requires service manipulation")
}

// Helper function
func getNextLesson(userID, skill string) (*NextLessonResponse, error) {
	req := NextLessonRequest{
		UserID:       userID,
		CurrentSkill: skill,
	}

	reqBody, _ := json.Marshal(req)
	resp, err := http.Post(
		AdaptiveEngineURL+"/api/adaptive/next-lesson",
		"application/json",
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var nextLesson NextLessonResponse
	if err := json.Unmarshal(body, &nextLesson); err != nil {
		return nil, err
	}

	return &nextLesson, nil
}
