package http

import "adaptive-engine/internal/adaptive"

// NextLessonRequest is the HTTP request for next lesson recommendation
type NextLessonRequest struct {
	UserID       string `json:"user_id" binding:"required"`
	CurrentSkill string `json:"current_skill" binding:"required"`
}

// ToRecommendInput converts HTTP request to UseCase input
func (r NextLessonRequest) ToRecommendInput() adaptive.RecommendInput {
	return adaptive.RecommendInput{
		UserID:       r.UserID,
		CurrentSkill: r.CurrentSkill,
	}
}
