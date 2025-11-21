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
