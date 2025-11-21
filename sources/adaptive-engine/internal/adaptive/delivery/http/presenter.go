package http

import "adaptive-engine/internal/adaptive"

// NextLessonResponse is the HTTP response for next lesson recommendation
type NextLessonResponse struct {
	NextLessonID int    `json:"next_lesson_id"`
	Reason       string `json:"reason"`
	MasteryScore int    `json:"mastery_score"`
	ContentType  string `json:"content_type"`
}

// ToNextLessonResponse converts UseCase output to HTTP response
func toResponse(o adaptive.RecommendOutput) NextLessonResponse {
	return NextLessonResponse{
		NextLessonID: o.NextLessonID,
		Reason:       o.Reason,
		MasteryScore: o.MasteryScore,
		ContentType:  o.ContentType,
	}
}
