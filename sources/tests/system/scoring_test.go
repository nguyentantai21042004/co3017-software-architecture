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

const (
	ScoringServiceURL = "http://localhost:8082"
	ContentServiceURL = "http://localhost:8081"
)

// SubmitAnswerRequest represents the request to submit an answer
type SubmitAnswerRequest struct {
	UserID     string `json:"user_id"`
	QuestionID int64  `json:"question_id"`
	Answer     string `json:"answer"`
}

// SubmitAnswerResponse represents the response from scoring service
type SubmitAnswerResponse struct {
	Correct bool   `json:"correct"`
	Score   int    `json:"score"`
	Message string `json:"message"`
}

// TC_SCORING_01: Submit Correct Answer & Publish Event
func TestScoringCorrectAnswer(t *testing.T) {
	fmt.Println("\n🧪 TC_SCORING_01: Submit Correct Answer & Publish Event")

	// Get a question first to know the correct answer
	question, err := getQuestion(6) // Math remedial question: "What is 2 + 2?"
	require.NoError(t, err, "Failed to get question")

	// Submit correct answer
	req := SubmitAnswerRequest{
		UserID:     fmt.Sprintf("test-scoring-01-%d", time.Now().UnixNano()),
		QuestionID: 6,
		Answer:     question.CorrectAnswer, // "4"
	}

	resp, err := submitAnswer(req)
	require.NoError(t, err, "Failed to submit answer")

	// Verify response
	assert.True(t, resp.Correct, "Answer should be marked as correct")
	assert.Equal(t, 100, resp.Score, "Score should be 100 for correct answer")

	fmt.Printf("✅ Response: correct=%v, score=%d\n", resp.Correct, resp.Score)

	// Wait for async event processing
	time.Sleep(2 * time.Second)

	// Verify event was published and processed by checking mastery update
	mastery, err := getMastery(req.UserID, "math")
	require.NoError(t, err, "Failed to get mastery")

	fmt.Printf("✅ Mastery updated: %d%%\n", mastery.MasteryScore)
	assert.Greater(t, mastery.MasteryScore, 0, "Mastery should be updated after correct answer")
}

// TC_SCORING_02: Submit Incorrect Answer & Publish Event
func TestScoringIncorrectAnswer(t *testing.T) {
	fmt.Println("\n🧪 TC_SCORING_02: Submit Incorrect Answer & Publish Event")

	// Submit incorrect answer
	req := SubmitAnswerRequest{
		UserID:     fmt.Sprintf("test-scoring-02-%d", time.Now().UnixNano()),
		QuestionID: 6,
		Answer:     "WRONG_ANSWER",
	}

	resp, err := submitAnswer(req)
	require.NoError(t, err, "Failed to submit answer")

	// Verify response
	assert.False(t, resp.Correct, "Answer should be marked as incorrect")
	assert.Equal(t, 0, resp.Score, "Score should be 0 for incorrect answer")

	fmt.Printf("✅ Response: correct=%v, score=%d\n", resp.Correct, resp.Score)

	// Wait for async event processing
	time.Sleep(2 * time.Second)

	// Verify event was still published (mastery record should exist even with 0 score)
	mastery, err := getMastery(req.UserID, "math")
	require.NoError(t, err, "Failed to get mastery - event should still be published")

	fmt.Printf("✅ Mastery after incorrect: %d%%\n", mastery.MasteryScore)
}

// TC_SCORING_03: Invalid Question ID
func TestScoringInvalidQuestionID(t *testing.T) {
	fmt.Println("\n🧪 TC_SCORING_03: Invalid Question ID")

	// Submit with non-existent question ID
	req := SubmitAnswerRequest{
		UserID:     fmt.Sprintf("test-scoring-03-%d", time.Now().UnixNano()),
		QuestionID: 99999, // Non-existent
		Answer:     "A",
	}

	// Make HTTP request
	reqBody, _ := json.Marshal(req)
	httpResp, err := http.Post(
		ScoringServiceURL+"/api/scoring/submit",
		"application/json",
		bytes.NewBuffer(reqBody),
	)
	require.NoError(t, err, "HTTP request should succeed")
	defer httpResp.Body.Close()

	// Verify error response
	assert.True(t, httpResp.StatusCode == 404 || httpResp.StatusCode == 400,
		"Should return 404 or 400 for invalid question ID, got %d", httpResp.StatusCode)

	fmt.Printf("✅ Correctly returned status code: %d\n", httpResp.StatusCode)

	// Verify no event was published (no mastery record should be created)
	time.Sleep(1 * time.Second)
	_, err = getMastery(req.UserID, "math")
	require.Error(t, err, "Should return error for invalid question ID")

	// It's OK if this returns an error or 0 mastery - the point is no submission was recorded
	fmt.Println("✅ No submission recorded for invalid question")
}

// TC_SCORING_04: RabbitMQ Downtime Resilience
func TestScoringRabbitMQResilience(t *testing.T) {
	fmt.Println("\n🧪 TC_SCORING_04: RabbitMQ Downtime Resilience")
	fmt.Println("⚠️  This test requires manual RabbitMQ stop/start")
	fmt.Println("⚠️  Skipping automated execution - verify manually:")
	fmt.Println("   1. Stop RabbitMQ: docker stop rabbitmq")
	fmt.Println("   2. Submit answer via API")
	fmt.Println("   3. Should return 200 (availability prioritized)")
	fmt.Println("   4. Check logs for 'Failed to publish event'")
	fmt.Println("   5. Restart RabbitMQ: docker start rabbitmq")

	t.Skip("Manual test - requires RabbitMQ manipulation")
}

// Helper functions

type QuestionResponse struct {
	ID            int64  `json:"id"`
	Content       string `json:"content"`
	CorrectAnswer string `json:"correct_answer"`
	SkillTag      string `json:"skill_tag"`
	IsRemedial    bool   `json:"is_remedial"`
}

type QuestionAPIResponse struct {
	Success bool             `json:"success"`
	Data    QuestionResponse `json:"data"`
}

func getQuestion(id int64) (*QuestionResponse, error) {
	resp, err := http.Get(fmt.Sprintf("%s/api/content/%d", ContentServiceURL, id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var apiResp QuestionAPIResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, err
	}

	return &apiResp.Data, nil
}

func submitAnswer(req SubmitAnswerRequest) (*SubmitAnswerResponse, error) {
	reqBody, _ := json.Marshal(req)
	resp, err := http.Post(
		ScoringServiceURL+"/api/scoring/submit",
		"application/json",
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var submitResp SubmitAnswerResponse
	if err := json.Unmarshal(body, &submitResp); err != nil {
		return nil, err
	}

	return &submitResp, nil
}

type MasteryResponse struct {
	MasteryScore int `json:"mastery_score"`
}

type MasteryResponseWrapper struct {
	ErrorCode int              `json:"error_code"`
	Message   string           `json:"message"`
	Data      MasteryResponse  `json:"data"`
}

func getMastery(userID, skill string) (*MasteryResponse, error) {
	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/internal/learner/%s/mastery?skill=%s", userID, skill))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var wrapper MasteryResponseWrapper
	if err := json.Unmarshal(body, &wrapper); err != nil {
		return nil, err
	}

	return &wrapper.Data, nil
}
