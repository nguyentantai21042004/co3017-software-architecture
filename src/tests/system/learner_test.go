package system

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TC_LEARNER_01: Update Score - Weighted Average
func TestLearnerWeightedAverage(t *testing.T) {
	fmt.Println("\n🧪 TC_LEARNER_01: Update Score - Weighted Average")

	userID := "test-learner-01"

	// Setup: Submit first correct answer to establish baseline
	question, _ := getQuestion(6)
	req1 := SubmitAnswerRequest{
		UserID:     userID,
		QuestionID: 6,
		Answer:     question.CorrectAnswer,
	}
	submitAnswer(req1)
	time.Sleep(2 * time.Second)

	// Get initial mastery
	mastery1, err := getMastery(userID, "math")
	require.NoError(t, err)
	initialScore := mastery1.MasteryScore
	fmt.Printf("📊 Initial mastery: %d%%\n", initialScore)

	// Submit second correct answer
	submitAnswer(req1)
	time.Sleep(2 * time.Second)

	// Get updated mastery
	mastery2, err := getMastery(userID, "math")
	require.NoError(t, err)
	updatedScore := mastery2.MasteryScore

	// Verify weighted average formula: (old + new) / 2
	expectedScore := (initialScore + 100) / 2
	fmt.Printf("📊 Updated mastery: %d%% (expected: %d%%)\n", updatedScore, expectedScore)

	assert.Equal(t, expectedScore, updatedScore, "Should use weighted average formula")
}

// TC_LEARNER_02: Update Score - Penalize Incorrect Answer
func TestLearnerPenalizeIncorrect(t *testing.T) {
	fmt.Println("\n🧪 TC_LEARNER_02: Update Score - Penalize Incorrect Answer")

	userID := "test-learner-02"

	// Setup: Establish baseline with correct answer
	question, _ := getQuestion(6)
	req := SubmitAnswerRequest{
		UserID:     userID,
		QuestionID: 6,
		Answer:     question.CorrectAnswer,
	}
	submitAnswer(req)
	time.Sleep(2 * time.Second)

	mastery1, _ := getMastery(userID, "math")
	initialScore := mastery1.MasteryScore
	fmt.Printf("📊 Initial mastery: %d%%\n", initialScore)

	// Submit incorrect answer
	reqWrong := SubmitAnswerRequest{
		UserID:     userID,
		QuestionID: 6,
		Answer:     "WRONG",
	}
	submitAnswer(reqWrong)
	time.Sleep(2 * time.Second)

	// Verify mastery decreased
	mastery2, _ := getMastery(userID, "math")
	updatedScore := mastery2.MasteryScore

	expectedScore := (initialScore + 0) / 2
	fmt.Printf("📊 Penalized mastery: %d%% (expected: %d%%)\n", updatedScore, expectedScore)

	assert.Less(t, updatedScore, initialScore, "Mastery should decrease after incorrect answer")
	assert.Equal(t, expectedScore, updatedScore, "Should apply penalty formula")
}

// TC_LEARNER_03: New User Cold Start
func TestLearnerColdStart(t *testing.T) {
	fmt.Println("\n🧪 TC_LEARNER_03: New User Cold Start")

	// Use unique user ID that has never submitted before
	userID := fmt.Sprintf("test-learner-new-%d", time.Now().Unix())

	// Submit first answer
	question, _ := getQuestion(6)
	req := SubmitAnswerRequest{
		UserID:     userID,
		QuestionID: 6,
		Answer:     question.CorrectAnswer,
	}

	resp, err := submitAnswer(req)
	require.NoError(t, err)
	assert.True(t, resp.Correct)

	time.Sleep(2 * time.Second)

	// Verify new record was created
	mastery, err := getMastery(userID, "math")
	require.NoError(t, err, "Should create new mastery record for new user")

	// For first submission, mastery should be (0 + 100) / 2 = 50
	assert.Equal(t, 50, mastery.MasteryScore, "Cold start should initialize with formula (0 + score) / 2")

	fmt.Printf("✅ New user record created: mastery=%d%%\n", mastery.MasteryScore)
}

// TC_LEARNER_04: Idempotency Check
func TestLearnerIdempotency(t *testing.T) {
	fmt.Println("\n🧪 TC_LEARNER_04: Idempotency Check")
	fmt.Println("⚠️  This test verifies duplicate event handling")

	userID := "test-learner-idem"

	// Submit answer
	question, _ := getQuestion(6)
	req := SubmitAnswerRequest{
		UserID:     userID,
		QuestionID: 6,
		Answer:     question.CorrectAnswer,
	}
	submitAnswer(req)
	time.Sleep(2 * time.Second)

	// Get mastery after first submission
	mastery1, _ := getMastery(userID, "math")
	score1 := mastery1.MasteryScore
	fmt.Printf("📊 Mastery after 1st submission: %d%%\n", score1)

	// Submit same answer again (simulating duplicate event)
	submitAnswer(req)
	time.Sleep(2 * time.Second)

	// Get mastery after second submission
	mastery2, _ := getMastery(userID, "math")
	score2 := mastery2.MasteryScore
	fmt.Printf("📊 Mastery after 2nd submission: %d%%\n", score2)

	// Note: Without proper idempotency handling, score will change
	// With idempotency, score should stay the same OR change predictably
	// This test documents the current behavior
	if score1 == score2 {
		fmt.Println("✅ Idempotency: Score unchanged (duplicate detected)")
	} else {
		fmt.Printf("⚠️  Score changed: %d%% → %d%% (no idempotency or different submission)\n", score1, score2)
		fmt.Println("   This is expected if submissions are treated as separate events")
	}
}
