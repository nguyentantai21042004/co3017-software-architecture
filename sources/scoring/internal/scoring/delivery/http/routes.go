package http

import "github.com/gin-gonic/gin"

// MapScoringRoutes maps routes for scoring module
func MapScoringRoutes(r *gin.RouterGroup, h Handler) {
	r.POST("/submit", h.SubmitAnswer)
	r.GET("/answered-questions", h.GetAnsweredQuestions)
}
