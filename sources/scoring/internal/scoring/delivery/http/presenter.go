package http

import "scoring/internal/scoring"

// SubmitResponse is the HTTP response for answer submission
type SubmitResponse struct {
	Correct  bool   `json:"correct"`
	Score    int    `json:"score"`
	Feedback string `json:"feedback"`
}

// ToSubmitResponse converts UseCase output to HTTP response
func ToSubmitResponse(output scoring.SubmitOutput) SubmitResponse {
	return SubmitResponse{
		Correct:  output.Correct,
		Score:    output.Score,
		Feedback: output.Feedback,
	}
}
