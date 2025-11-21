package http

import "github.com/gin-gonic/gin"

// MapAdaptiveRoutes maps routes for adaptive module
func MapAdaptiveRoutes(r *gin.RouterGroup, h Handler) {
	r.POST("/next-lesson", func(c *gin.Context) {
		h.NextLesson(c)
	})
}
