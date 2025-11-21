package integration

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestFullLearningProgression tests complete progression from 0% to high mastery
func TestFullLearningProgression(t *testing.T) {
	userID := "test-user-progression"
	skill := SkillMath

	fmt.Printf("\n🧪 Test: Full Learning Progression (0%% → High Mastery)\n")
	fmt.Printf("User: %s, Skill: %s\n\n", userID, skill)

	// Track mastery progression
	masteryHistory := []int{}

	// Get initial mastery
	initialMastery, err := GetMastery(userID, skill)
	require.NoError(t, err)
	masteryHistory = append(masteryHistory, initialMastery.MasteryScore)
	fmt.Printf("📊 Initial mastery: %d%%\n\n", initialMastery.MasteryScore)

	// Submit 5 correct answers to build mastery
	for i := 1; i <= 5; i++ {
		fmt.Printf("--- Submission %d ---\n", i)

		// Get lesson
		lesson, err := GetNextLesson(userID, skill)
		require.NoError(t, err)
		fmt.Printf("Question ID: %d, Type: %s\n", lesson.NextLessonID, lesson.ContentType)

		// Verify content type based on mastery
		currentMastery := masteryHistory[len(masteryHistory)-1]
		if currentMastery < 50 {
			assert.Equal(t, "remedial", lesson.ContentType, "Should recommend remedial for mastery < 50")
		} else {
			assert.Equal(t, "standard", lesson.ContentType, "Should recommend standard for mastery >= 50")
		}

		// Get question and submit correct answer
		question, err := GetQuestion(int64(lesson.NextLessonID))
		require.NoError(t, err)

		submitResp, err := SubmitAnswer(userID, int64(lesson.NextLessonID), question.CorrectAnswer)
		require.NoError(t, err)
		assert.True(t, submitResp.Correct)
		fmt.Printf("Score: %d\n", submitResp.Score)

		// Calculate expected mastery
		expectedMastery := (currentMastery + submitResp.Score) / 2

		// Wait for update
		updatedMastery, err := WaitForMasteryUpdate(userID, skill, expectedMastery, MasteryUpdateTimeout)
		require.NoError(t, err)
		masteryHistory = append(masteryHistory, updatedMastery.MasteryScore)
		fmt.Printf("Mastery: %d%% → %d%%\n\n", currentMastery, updatedMastery.MasteryScore)
	}

	// Verify progression
	fmt.Println("📊 Mastery Progression:")
	for i, mastery := range masteryHistory {
		if i == 0 {
			fmt.Printf("   Initial: %d%%\n", mastery)
		} else {
			fmt.Printf("   After submission %d: %d%%\n", i, mastery)
		}
	}

	// Verify mastery increased
	finalMastery := masteryHistory[len(masteryHistory)-1]
	assert.Greater(t, finalMastery, initialMastery.MasteryScore, "Mastery should increase with correct answers")

	fmt.Println("\n✅ Test completed successfully!")
	fmt.Printf("Final mastery: %d%%\n", finalMastery)
}

// TestMixedCorrectIncorrectAnswers tests mastery changes with mixed results
func TestMixedCorrectIncorrectAnswers(t *testing.T) {
	userID := "test-user-mixed"
	skill := SkillMath

	fmt.Printf("\n🧪 Test: Mixed Correct/Incorrect Answers\n")
	fmt.Printf("User: %s, Skill: %s\n\n", userID, skill)

	// Pattern: Correct, Correct, Incorrect, Correct, Incorrect
	answerPattern := []bool{true, true, false, true, false}
	masteryHistory := []int{}

	// Get initial mastery
	initialMastery, err := GetMastery(userID, skill)
	require.NoError(t, err)
	masteryHistory = append(masteryHistory, initialMastery.MasteryScore)
	fmt.Printf("📊 Initial mastery: %d%%\n\n", initialMastery.MasteryScore)

	for i, shouldBeCorrect := range answerPattern {
		fmt.Printf("--- Submission %d (%s) ---\n", i+1, map[bool]string{true: "Correct", false: "Incorrect"}[shouldBeCorrect])

		// Get lesson
		lesson, err := GetNextLesson(userID, skill)
		require.NoError(t, err)

		// Get question
		question, err := GetQuestion(int64(lesson.NextLessonID))
		require.NoError(t, err)

		// Submit answer (correct or incorrect based on pattern)
		var answer string
		if shouldBeCorrect {
			answer = question.CorrectAnswer
		} else {
			answer = "WRONG_ANSWER"
		}

		submitResp, err := SubmitAnswer(userID, int64(lesson.NextLessonID), answer)
		require.NoError(t, err)
		assert.Equal(t, shouldBeCorrect, submitResp.Correct)
		fmt.Printf("Score: %d\n", submitResp.Score)

		// Calculate expected mastery
		currentMastery := masteryHistory[len(masteryHistory)-1]
		expectedMastery := (currentMastery + submitResp.Score) / 2

		// Wait for update
		updatedMastery, err := WaitForMasteryUpdate(userID, skill, expectedMastery, MasteryUpdateTimeout)
		require.NoError(t, err)
		masteryHistory = append(masteryHistory, updatedMastery.MasteryScore)
		fmt.Printf("Mastery: %d%% → %d%%\n\n", currentMastery, updatedMastery.MasteryScore)
	}

	// Verify progression
	fmt.Println("📊 Mastery Progression with Mixed Answers:")
	for i, mastery := range masteryHistory {
		if i == 0 {
			fmt.Printf("   Initial: %d%%\n", mastery)
		} else {
			result := "✓"
			if !answerPattern[i-1] {
				result = "✗"
			}
			fmt.Printf("   After submission %d (%s): %d%%\n", i, result, mastery)
		}
	}

	fmt.Println("\n✅ Test completed successfully!")
}

// TestMultipleSkills tests learning across different skills
func TestMultipleSkills(t *testing.T) {
	userID := "test-user-multiskill"

	fmt.Printf("\n🧪 Test: Multiple Skills Learning\n")
	fmt.Printf("User: %s\n\n", userID)

	skills := []string{SkillMath, SkillScience}
	skillMastery := make(map[string][]int)

	// Initialize mastery tracking for each skill
	for _, skill := range skills {
		mastery, err := GetMastery(userID, skill)
		require.NoError(t, err)
		skillMastery[skill] = []int{mastery.MasteryScore}
		fmt.Printf("📊 Initial %s mastery: %d%%\n", skill, mastery.MasteryScore)
	}
	fmt.Println()

	// Submit 2 correct answers for each skill
	for _, skill := range skills {
		fmt.Printf("=== Learning %s ===\n", skill)

		for i := 1; i <= 2; i++ {
			fmt.Printf("Submission %d:\n", i)

			// Get lesson
			lesson, err := GetNextLesson(userID, skill)
			require.NoError(t, err)

			// Get question and submit correct answer
			question, err := GetQuestion(int64(lesson.NextLessonID))
			require.NoError(t, err)

			submitResp, err := SubmitAnswer(userID, int64(lesson.NextLessonID), question.CorrectAnswer)
			require.NoError(t, err)
			assert.True(t, submitResp.Correct)

			// Calculate expected mastery
			currentMastery := skillMastery[skill][len(skillMastery[skill])-1]
			expectedMastery := (currentMastery + 100) / 2

			// Wait for update
			updatedMastery, err := WaitForMasteryUpdate(userID, skill, expectedMastery, MasteryUpdateTimeout)
			require.NoError(t, err)
			skillMastery[skill] = append(skillMastery[skill], updatedMastery.MasteryScore)
			fmt.Printf("  %s mastery: %d%% → %d%%\n", skill, currentMastery, updatedMastery.MasteryScore)
		}
		fmt.Println()
	}

	// Verify each skill progressed independently
	fmt.Println("📊 Final Mastery by Skill:")
	for _, skill := range skills {
		history := skillMastery[skill]
		finalMastery := history[len(history)-1]
		initialMastery := history[0]
		fmt.Printf("   %s: %d%% → %d%%\n", skill, initialMastery, finalMastery)
		assert.Greater(t, finalMastery, initialMastery, "Mastery should increase for "+skill)
	}

	fmt.Println("\n✅ Test completed successfully!")
}

// TestBoundaryMasteryScores tests edge cases at mastery boundaries
func TestBoundaryMasteryScores(t *testing.T) {
	userID := "test-user-boundary"
	skill := SkillMath

	fmt.Printf("\n🧪 Test: Boundary Mastery Scores\n")
	fmt.Printf("User: %s, Skill: %s\n\n", userID, skill)

	// Build up to exactly 50% mastery (boundary)
	fmt.Println("📊 Building mastery to 50% (boundary)...")

	currentMastery := 0
	submissionCount := 0

	for currentMastery < 50 {
		submissionCount++

		// Get lesson
		lesson, err := GetNextLesson(userID, skill)
		require.NoError(t, err)

		// Should be remedial until we hit 50%
		if currentMastery < 50 {
			assert.Equal(t, "remedial", lesson.ContentType, "Should recommend remedial below 50%%")
		}

		// Get question and submit correct answer
		question, err := GetQuestion(int64(lesson.NextLessonID))
		require.NoError(t, err)

		_, err = SubmitAnswer(userID, int64(lesson.NextLessonID), question.CorrectAnswer)
		require.NoError(t, err)

		// Calculate expected mastery
		expectedMastery := (currentMastery + 100) / 2

		// Wait for update
		updatedMastery, err := WaitForMasteryUpdate(userID, skill, expectedMastery, MasteryUpdateTimeout)
		require.NoError(t, err)

		fmt.Printf("  Submission %d: %d%% → %d%%\n", submissionCount, currentMastery, updatedMastery.MasteryScore)
		currentMastery = updatedMastery.MasteryScore

		// Break if we've reached or exceeded 50%
		if currentMastery >= 50 {
			break
		}
	}

	fmt.Printf("\n✅ Reached boundary: %d%%\n", currentMastery)

	// Now verify that next lesson is standard
	fmt.Println("\n📚 Testing content type at boundary...")
	lesson, err := GetNextLesson(userID, skill)
	require.NoError(t, err)

	if currentMastery >= 50 {
		assert.Equal(t, "standard", lesson.ContentType, "Should recommend standard at/above 50%%")
		fmt.Printf("✓ Correctly recommending %s content at %d%% mastery\n", lesson.ContentType, currentMastery)
	}

	fmt.Println("\n✅ Test completed successfully!")
}

// TestRapidSubmissions tests system behavior with rapid consecutive submissions
func TestRapidSubmissions(t *testing.T) {
	userID := "test-user-rapid"
	skill := SkillScience

	fmt.Printf("\n🧪 Test: Rapid Consecutive Submissions\n")
	fmt.Printf("User: %s, Skill: %s\n\n", userID, skill)

	// Get initial mastery
	initialMastery, err := GetMastery(userID, skill)
	require.NoError(t, err)
	fmt.Printf("📊 Initial mastery: %d%%\n\n", initialMastery.MasteryScore)

	// Submit 3 answers rapidly
	fmt.Println("⚡ Submitting 3 answers rapidly...")
	expectedMastery := initialMastery.MasteryScore

	for i := 1; i <= 3; i++ {
		// Get lesson
		lesson, err := GetNextLesson(userID, skill)
		require.NoError(t, err)

		// Get question and submit
		question, err := GetQuestion(int64(lesson.NextLessonID))
		require.NoError(t, err)

		_, err = SubmitAnswer(userID, int64(lesson.NextLessonID), question.CorrectAnswer)
		require.NoError(t, err)

		// Calculate expected mastery
		expectedMastery = (expectedMastery + 100) / 2
		fmt.Printf("  Submission %d: Expected mastery: %d%%\n", i, expectedMastery)

		// Small delay to allow async processing
		time.Sleep(500 * time.Millisecond)
	}

	// Wait for final mastery to settle
	fmt.Println("\n⏳ Waiting for all updates to complete...")
	time.Sleep(3 * time.Second)

	// Verify final mastery
	finalMastery, err := GetMastery(userID, skill)
	require.NoError(t, err)

	fmt.Printf("\n📊 Final mastery: %d%% (expected: %d%%)\n", finalMastery.MasteryScore, expectedMastery)
	assert.Equal(t, expectedMastery, finalMastery.MasteryScore, "Final mastery should match expected value")

	fmt.Println("\n✅ Test completed successfully!")
}
