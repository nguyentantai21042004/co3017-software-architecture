package http

import "github.com/gin-gonic/gin"

// MapLearnerRoutes maps routes for learner module
func MapLearnerRoutes(r *gin.RouterGroup, h Handler) {
	r.GET("/:user_id/mastery", h.GetMastery)
}
