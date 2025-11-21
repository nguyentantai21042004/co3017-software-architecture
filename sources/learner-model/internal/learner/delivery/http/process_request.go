package http

import "learner-model-service/internal/learner"

// GetMasteryRequest represents the request parameters for getting mastery
type GetMasteryRequest struct {
	UserID   string
	SkillTag string
}

// ToGetMasteryInput converts HTTP request to UseCase input
func (r GetMasteryRequest) ToGetMasteryInput() learner.GetMasteryInput {
	return learner.GetMasteryInput{
		UserID:   r.UserID,
		SkillTag: r.SkillTag,
	}
}

// UpdateMasteryRequest represents the request body for updating mastery
type UpdateMasteryRequest struct {
	UserID        string `json:"user_id"`
	SkillTag      string `json:"skill_tag"`
	ScoreObtained int    `json:"score_obtained"`
}

// ToUpdateMasteryInput converts HTTP request to UseCase input
func (r UpdateMasteryRequest) ToUpdateMasteryInput() learner.UpdateMasteryInput {
	return learner.UpdateMasteryInput{
		UserID:        r.UserID,
		SkillTag:      r.SkillTag,
		ScoreObtained: r.ScoreObtained,
	}
}
