package integration

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestCompleteLearningFlow_LowMasteryToHigh tests the complete adaptive learning flow
// for a user starting with low mastery and progressing to high mastery
func TestCompleteLearningFlow_LowMasteryToHigh(t *testing.T) {
	userID := TestUser1
	skill := SkillMath

	fmt.Printf("\n🧪 Test: Complete Learning Flow - Low Mastery to High\n")
	fmt.Printf("User: %s, Skill: %s\n\n", userID, skill)

	// Step 1: Get initial mastery (should be 0 for new user)
	fmt.Println("📊 Step 1: Check initial mastery...")
	initialMastery, err := GetMastery(userID, skill)
	require.NoError(t, err, "Failed to get initial mastery")
	fmt.Printf("   Initial mastery: %d%%\n", initialMastery.MasteryScore)

	// Step 2: Get first lesson recommendation (should be remedial since mastery is low)
	fmt.Println("\n📚 Step 2: Get first lesson recommendation...")
	lesson1, err := GetNextLesson(userID, skill)
	require.NoError(t, err, "Failed to get first lesson")
	fmt.Printf("   Recommended: Question ID %d, Type: %s, Reason: %s\n",
		lesson1.NextLessonID, lesson1.ContentType, lesson1.Reason)

	if initialMastery.MasteryScore < 50 {
		assert.Equal(t, "remedial", lesson1.ContentType, "Should recommend remedial content for low mastery")
	}

	// Step 3: Get question details to know the correct answer
	fmt.Println("\n❓ Step 3: Fetch question details...")
	question1, err := GetQuestion(int64(lesson1.NextLessonID))
	require.NoError(t, err, "Failed to get question details")
	fmt.Printf("   Question: %s\n", question1.Content)
	fmt.Printf("   Correct Answer: %s\n", question1.CorrectAnswer)

	// Step 4: Submit correct answer
	fmt.Println("\n✍️  Step 4: Submit correct answer...")
	submitResp1, err := SubmitAnswer(userID, int64(lesson1.NextLessonID), question1.CorrectAnswer)
	require.NoError(t, err, "Failed to submit answer")
	assert.True(t, submitResp1.Correct, "Answer should be marked as correct")
	assert.Equal(t, 100, submitResp1.Score, "Correct answer should get 100 points")
	fmt.Printf("   Result: %s (Score: %d)\n", submitResp1.Feedback, submitResp1.Score)

	// Step 5: Wait for mastery update via RabbitMQ
	// Expected: (0 + 100) / 2 = 50
	fmt.Println("\n⏳ Step 5: Waiting for mastery update (async via RabbitMQ)...")
	expectedMastery1 := (initialMastery.MasteryScore + 100) / 2
	updatedMastery1, err := WaitForMasteryUpdate(userID, skill, expectedMastery1, MasteryUpdateTimeout)
	require.NoError(t, err, "Mastery should be updated within timeout")
	assert.Equal(t, expectedMastery1, updatedMastery1.MasteryScore, "Mastery should be calculated correctly")
	fmt.Printf("   ✅ Mastery updated: %d%% → %d%%\n", initialMastery.MasteryScore, updatedMastery1.MasteryScore)

	// Step 6: Get second lesson recommendation (should now be standard if mastery >= 50)
	fmt.Println("\n📚 Step 6: Get second lesson recommendation...")
	lesson2, err := GetNextLesson(userID, skill)
	require.NoError(t, err, "Failed to get second lesson")
	fmt.Printf("   Recommended: Question ID %d, Type: %s, Reason: %s\n",
		lesson2.NextLessonID, lesson2.ContentType, lesson2.Reason)

	if updatedMastery1.MasteryScore >= 50 {
		assert.Equal(t, "standard", lesson2.ContentType, "Should recommend standard content for mastery >= 50")
	}

	// Step 7: Submit another correct answer
	fmt.Println("\n✍️  Step 7: Submit second correct answer...")
	question2, err := GetQuestion(int64(lesson2.NextLessonID))
	require.NoError(t, err, "Failed to get second question")

	submitResp2, err := SubmitAnswer(userID, int64(lesson2.NextLessonID), question2.CorrectAnswer)
	require.NoError(t, err, "Failed to submit second answer")
	assert.True(t, submitResp2.Correct, "Second answer should be correct")
	fmt.Printf("   Result: %s (Score: %d)\n", submitResp2.Feedback, submitResp2.Score)

	// Step 8: Wait for second mastery update
	// Expected: (50 + 100) / 2 = 75
	fmt.Println("\n⏳ Step 8: Waiting for second mastery update...")
	expectedMastery2 := (updatedMastery1.MasteryScore + 100) / 2
	updatedMastery2, err := WaitForMasteryUpdate(userID, skill, expectedMastery2, MasteryUpdateTimeout)
	require.NoError(t, err, "Second mastery update should complete")
	assert.Equal(t, expectedMastery2, updatedMastery2.MasteryScore, "Second mastery should be calculated correctly")
	fmt.Printf("   ✅ Mastery updated: %d%% → %d%%\n", updatedMastery1.MasteryScore, updatedMastery2.MasteryScore)

	fmt.Println("\n✅ Test completed successfully!")
	fmt.Printf("Final mastery: %d%%\n", updatedMastery2.MasteryScore)
}

// TestCompleteLearningFlow_HighMasteryToLow tests mastery degradation
// when a user with high mastery submits incorrect answers
func TestCompleteLearningFlow_HighMasteryToLow(t *testing.T) {
	userID := TestUser2
	skill := SkillScience

	fmt.Printf("\n🧪 Test: Complete Learning Flow - High Mastery to Low\n")
	fmt.Printf("User: %s, Skill: %s\n\n", userID, skill)

	// Note: This test assumes the user already has high mastery (e.g., 80%)
	// In a real scenario, you'd need to set this up first or use a pre-seeded user

	// Step 1: Check initial mastery
	fmt.Println("📊 Step 1: Check initial mastery...")
	initialMastery, err := GetMastery(userID, skill)
	require.NoError(t, err, "Failed to get initial mastery")
	fmt.Printf("   Initial mastery: %d%%\n", initialMastery.MasteryScore)

	// If mastery is 0, we need to build it up first
	if initialMastery.MasteryScore == 0 {
		fmt.Println("   ⚠️  User has 0 mastery, building up to 80% first...")

		// Submit 4 correct answers to reach ~94% mastery
		for i := 0; i < 4; i++ {
			lesson, err := GetNextLesson(userID, skill)
			require.NoError(t, err)

			question, err := GetQuestion(int64(lesson.NextLessonID))
			require.NoError(t, err)

			_, err = SubmitAnswer(userID, int64(lesson.NextLessonID), question.CorrectAnswer)
			require.NoError(t, err)

			// Wait a bit for mastery to update
			time.Sleep(2 * time.Second)
		}

		initialMastery, err = GetMastery(userID, skill)
		require.NoError(t, err)
		fmt.Printf("   Built up mastery to: %d%%\n", initialMastery.MasteryScore)
	}

	// Step 2: Get lesson recommendation (should be standard for high mastery)
	fmt.Println("\n📚 Step 2: Get lesson recommendation...")
	lesson, err := GetNextLesson(userID, skill)
	require.NoError(t, err, "Failed to get lesson")
	fmt.Printf("   Recommended: Question ID %d, Type: %s\n", lesson.NextLessonID, lesson.ContentType)

	if initialMastery.MasteryScore >= 50 {
		assert.Equal(t, "standard", lesson.ContentType, "Should recommend standard content for high mastery")
	}

	// Step 3: Submit INCORRECT answer
	fmt.Println("\n✍️  Step 3: Submit incorrect answer...")
	question, err := GetQuestion(int64(lesson.NextLessonID))
	require.NoError(t, err)

	wrongAnswer := "WRONG_ANSWER"
	if question.CorrectAnswer == "WRONG_ANSWER" {
		wrongAnswer = "ANOTHER_WRONG_ANSWER"
	}

	submitResp, err := SubmitAnswer(userID, int64(lesson.NextLessonID), wrongAnswer)
	require.NoError(t, err, "Failed to submit answer")
	assert.False(t, submitResp.Correct, "Answer should be marked as incorrect")
	assert.Equal(t, 0, submitResp.Score, "Incorrect answer should get 0 points")
	fmt.Printf("   Result: %s (Score: %d)\n", submitResp.Feedback, submitResp.Score)

	// Step 4: Wait for mastery update (should decrease)
	// Expected: (initialMastery + 0) / 2
	fmt.Println("\n⏳ Step 4: Waiting for mastery update...")
	expectedMastery := (initialMastery.MasteryScore + 0) / 2
	updatedMastery, err := WaitForMasteryUpdate(userID, skill, expectedMastery, MasteryUpdateTimeout)
	require.NoError(t, err, "Mastery should be updated")
	assert.Equal(t, expectedMastery, updatedMastery.MasteryScore, "Mastery should decrease")
	fmt.Printf("   ✅ Mastery updated: %d%% → %d%%\n", initialMastery.MasteryScore, updatedMastery.MasteryScore)

	// Step 5: Get next lesson (should switch to remedial if mastery dropped below 50)
	fmt.Println("\n📚 Step 5: Get next lesson after mastery drop...")
	nextLesson, err := GetNextLesson(userID, skill)
	require.NoError(t, err, "Failed to get next lesson")
	fmt.Printf("   Recommended: Question ID %d, Type: %s\n", nextLesson.NextLessonID, nextLesson.ContentType)

	if updatedMastery.MasteryScore < 50 {
		assert.Equal(t, "remedial", nextLesson.ContentType, "Should switch to remedial content for low mastery")
	}

	fmt.Println("\n✅ Test completed successfully!")
	fmt.Printf("Final mastery: %d%%\n", updatedMastery.MasteryScore)
}

// TestMultipleSubmissions tests mastery progression through multiple submissions
func TestMultipleSubmissions(t *testing.T) {
	userID := TestUser3
	skill := SkillMath

	fmt.Printf("\n🧪 Test: Multiple Submissions - Mastery Progression\n")
	fmt.Printf("User: %s, Skill: %s\n\n", userID, skill)

	// Track mastery progression
	masteryProgression := []int{}

	// Get initial mastery
	initialMastery, err := GetMastery(userID, skill)
	require.NoError(t, err)
	masteryProgression = append(masteryProgression, initialMastery.MasteryScore)
	fmt.Printf("Starting mastery: %d%%\n\n", initialMastery.MasteryScore)

	// Submit 3 correct answers and track progression
	for i := 1; i <= 3; i++ {
		fmt.Printf("--- Submission %d ---\n", i)

		// Get lesson
		lesson, err := GetNextLesson(userID, skill)
		require.NoError(t, err)
		fmt.Printf("Question ID: %d, Type: %s\n", lesson.NextLessonID, lesson.ContentType)

		// Get question and submit correct answer
		question, err := GetQuestion(int64(lesson.NextLessonID))
		require.NoError(t, err)

		submitResp, err := SubmitAnswer(userID, int64(lesson.NextLessonID), question.CorrectAnswer)
		require.NoError(t, err)
		assert.True(t, submitResp.Correct)
		fmt.Printf("Score: %d\n", submitResp.Score)

		// Calculate expected mastery
		currentMastery := masteryProgression[len(masteryProgression)-1]
		expectedMastery := (currentMastery + submitResp.Score) / 2

		// Wait for update
		updatedMastery, err := WaitForMasteryUpdate(userID, skill, expectedMastery, MasteryUpdateTimeout)
		require.NoError(t, err)
		masteryProgression = append(masteryProgression, updatedMastery.MasteryScore)
		fmt.Printf("Mastery: %d%% → %d%%\n\n", currentMastery, updatedMastery.MasteryScore)
	}

	// Verify progression
	fmt.Println("📊 Mastery Progression:")
	for i, mastery := range masteryProgression {
		if i == 0 {
			fmt.Printf("   Initial: %d%%\n", mastery)
		} else {
			fmt.Printf("   After submission %d: %d%%\n", i, mastery)
		}
	}

	// Verify mastery is increasing (assuming all correct answers)
	for i := 1; i < len(masteryProgression); i++ {
		assert.GreaterOrEqual(t, masteryProgression[i], masteryProgression[i-1],
			"Mastery should increase or stay same with correct answers")
	}

	fmt.Println("\n✅ Test completed successfully!")
}
