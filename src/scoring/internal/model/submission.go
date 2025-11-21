package model

import "time"

// Submission represents a user's answer submission
type Submission struct {
	ID              int64     `json:"id" db:"id"`
	UserID          string    `json:"user_id" db:"user_id"`
	QuestionID      int64     `json:"question_id" db:"question_id"`
	SubmittedAnswer string    `json:"submitted_answer" db:"submitted_answer"`
	ScoreAwarded    int       `json:"score_awarded" db:"score_awarded"`
	IsPassed        bool      `json:"is_passed" db:"is_passed"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
}

// SubmitRequest is the request payload for submitting an answer
type SubmitRequest struct {
	UserID     string `json:"user_id" binding:"required"`
	QuestionID int64  `json:"question_id" binding:"required"`
	Answer     string `json:"answer" binding:"required"`
}

// SubmitResponse is the response after scoring
type SubmitResponse struct {
	Correct  bool   `json:"correct"`
	Score    int    `json:"score"`
	Feedback string `json:"feedback"`
}

// SubmissionEvent is the event published to RabbitMQ
type SubmissionEvent struct {
	Event         string `json:"event"`
	UserID        string `json:"user_id"`
	SkillTag      string `json:"skill_tag"`
	ScoreObtained int    `json:"score_obtained"`
	Timestamp     string `json:"timestamp"`
}
