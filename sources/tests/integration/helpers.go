package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// HTTP client with timeout
var httpClient = &http.Client{
	Timeout: HTTPTimeout,
}

// Request/Response models

type NextLessonRequest struct {
	UserID       string `json:"user_id"`
	CurrentSkill string `json:"current_skill"`
}

type NextLessonResponse struct {
	NextLessonID int    `json:"next_lesson_id"`
	Reason       string `json:"reason"`
	MasteryScore int    `json:"mastery_score"`
	ContentType  string `json:"content_type"`
}

type SubmitAnswerRequest struct {
	UserID     string `json:"user_id"`
	QuestionID int64  `json:"question_id"`
	Answer     string `json:"answer"`
}

type SubmitAnswerResponse struct {
	Correct  bool   `json:"correct"`
	Score    int    `json:"score"`
	Feedback string `json:"feedback"`
}

type MasteryResponse struct {
	UserID       string `json:"user_id"`
	SkillTag     string `json:"skill_tag"`
	MasteryScore int    `json:"mastery_score"`
	LastUpdated  string `json:"last_updated"`
}

type QuestionResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    QuestionData `json:"data"`
}

type QuestionData struct {
	ID            int64  `json:"id"`
	Content       string `json:"content"`
	CorrectAnswer string `json:"correct_answer"`
	SkillTag      string `json:"skill_tag"`
	IsRemedial    bool   `json:"is_remedial"`
}

// Helper functions

// GetNextLesson calls Adaptive Engine to get next lesson recommendation
func GetNextLesson(userID, skill string) (*NextLessonResponse, error) {
	req := NextLessonRequest{
		UserID:       userID,
		CurrentSkill: skill,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := httpClient.Post(
		fmt.Sprintf("%s/api/adaptive/next-lesson", AdaptiveEngineURL),
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("adaptive engine returned %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var result NextLessonResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// SubmitAnswer calls Scoring Service to submit an answer
func SubmitAnswer(userID string, questionID int64, answer string) (*SubmitAnswerResponse, error) {
	req := SubmitAnswerRequest{
		UserID:     userID,
		QuestionID: questionID,
		Answer:     answer,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := httpClient.Post(
		fmt.Sprintf("%s/api/scoring/submit", ScoringServiceURL),
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("scoring service returned %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var result SubmitAnswerResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetMastery calls Learner Model to get current mastery score
func GetMastery(userID, skill string) (*MasteryResponse, error) {
	url := fmt.Sprintf("%s/internal/learner/%s/mastery?skill=%s", LearnerModelURL, userID, skill)

	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("learner model returned %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var result MasteryResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetQuestion calls Content Service to get question details
func GetQuestion(questionID int64) (*QuestionData, error) {
	url := fmt.Sprintf("%s/api/content/%d", ContentServiceURL, questionID)

	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("content service returned %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var result QuestionResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if !result.Success {
		return nil, fmt.Errorf("content service error: %s", result.Message)
	}

	return &result.Data, nil
}

// WaitForMasteryUpdate polls Learner Model until mastery reaches expected score or timeout
func WaitForMasteryUpdate(userID, skill string, expectedScore int, timeout time.Duration) (*MasteryResponse, error) {
	deadline := time.Now().Add(timeout)

	for time.Now().Before(deadline) {
		mastery, err := GetMastery(userID, skill)
		if err != nil {
			time.Sleep(PollingInterval)
			continue
		}

		if mastery.MasteryScore == expectedScore {
			return mastery, nil
		}

		time.Sleep(PollingInterval)
	}

	// One final attempt
	mastery, err := GetMastery(userID, skill)
	if err != nil {
		return nil, fmt.Errorf("timeout waiting for mastery update: %w", err)
	}

	if mastery.MasteryScore != expectedScore {
		return nil, fmt.Errorf("mastery score mismatch: expected %d, got %d", expectedScore, mastery.MasteryScore)
	}

	return mastery, nil
}

// CleanupUserMastery removes test user mastery data (optional, for cleanup)
func CleanupUserMastery(userID, skill string) error {
	// This would require a DELETE endpoint in Learner Model Service
	// For now, we'll just log that cleanup is needed
	fmt.Printf("⚠️  Manual cleanup needed: DELETE mastery for user=%s, skill=%s\n", userID, skill)
	return nil
}
