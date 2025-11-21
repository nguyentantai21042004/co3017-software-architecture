package http

import "learner-model-service/internal/learner"

// MasteryResponse is the HTTP response for mastery query
type MasteryResponse struct {
	UserID       string `json:"user_id"`
	SkillTag     string `json:"skill_tag"`
	MasteryScore int    `json:"mastery_score"`
	LastUpdated  string `json:"last_updated"`
}

// ToMasteryResponse converts UseCase output to HTTP response
func ToMasteryResponse(output learner.MasteryOutput) MasteryResponse {
	return MasteryResponse{
		UserID:       output.UserID,
		SkillTag:     output.SkillTag,
		MasteryScore: output.MasteryScore,
		LastUpdated:  output.LastUpdated,
	}
}
