package model

// NextLessonRequest is the request for adaptive content recommendation
type NextLessonRequest struct {
	UserID      string `json:"user_id" binding:"required"`
	CurrentSkill string `json:"current_skill" binding:"required"`
}

// NextLessonResponse is the response with recommended content
type NextLessonResponse struct {
	NextLessonID int    `json:"next_lesson_id"`
	Reason       string `json:"reason"`
	MasteryScore int    `json:"mastery_score"`
	ContentType  string `json:"content_type"` // "remedial" or "standard"
}

// MasteryResponse from Learner Model Service
type MasteryResponse struct {
	UserID       string `json:"user_id"`
	SkillTag     string `json:"skill_tag"`
	MasteryScore int    `json:"mastery_score"`
	LastUpdated  string `json:"last_updated"`
}

// QuestionResponse from Content Service
type QuestionResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    QuestionData `json:"data"`
}

type QuestionData struct {
	ID             int64  `json:"id"`
	Content        string `json:"content"`
	CorrectAnswer  string `json:"correct_answer"`
	SkillTag       string `json:"skill_tag"`
	DifficultyLevel int   `json:"difficulty_level"`
	IsRemedial     bool   `json:"is_remedial"`
}
