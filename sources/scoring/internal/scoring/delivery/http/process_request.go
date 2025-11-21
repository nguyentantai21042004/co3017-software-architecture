package http

import "scoring/internal/scoring"

// SubmitRequest is the HTTP request for answer submission
type SubmitRequest struct {
	UserID     string `json:"user_id"`
	QuestionID int64  `json:"question_id"`
	Answer     string `json:"answer"`
}

// ToSubmitInput converts HTTP request to UseCase input
func (r SubmitRequest) ToSubmitInput() scoring.SubmitInput {
	return scoring.SubmitInput{
		UserID:     r.UserID,
		QuestionID: r.QuestionID,
		Answer:     r.Answer,
	}
}
