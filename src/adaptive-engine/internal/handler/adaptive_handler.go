package handler

import (
	"net/http"

	"adaptive-engine-service/internal/model"
	"adaptive-engine-service/internal/service"

	"github.com/gin-gonic/gin"
)

type AdaptiveHandler struct {
	service service.AdaptiveService
}

func NewAdaptiveHandler(service service.AdaptiveService) *AdaptiveHandler {
	return &AdaptiveHandler{service: service}
}

func (h *AdaptiveHandler) NextLesson(c *gin.Context) {
	var req model.NextLessonRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	response, err := h.service.RecommendNextLesson(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *AdaptiveHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"service": "adaptive-engine-service",
	})
}
