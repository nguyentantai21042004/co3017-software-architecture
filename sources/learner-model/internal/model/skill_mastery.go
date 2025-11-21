package model

import "time"

// SkillMastery represents a user's mastery level for a specific skill
type SkillMastery struct {
	UserID       string    `json:"user_id" db:"user_id"`
	SkillTag     string    `json:"skill_tag" db:"skill_tag"`
	CurrentScore int       `json:"current_score" db:"current_score"`
	LastUpdated  time.Time `json:"last_updated" db:"last_updated"`
}

// SubmissionEvent is the event received from RabbitMQ
type SubmissionEvent struct {
	Event         string `json:"event"`
	UserID        string `json:"user_id"`
	SkillTag      string `json:"skill_tag"`
	ScoreObtained int    `json:"score_obtained"`
	Timestamp     string `json:"timestamp"`
}
