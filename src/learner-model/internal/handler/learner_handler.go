package handler

import (
	"net/http"

	"learner-model-service/internal/service"

	"github.com/gin-gonic/gin"
)

type LearnerHandler struct {
	service service.LearnerService
}

func NewLearnerHandler(service service.LearnerService) *LearnerHandler {
	return &LearnerHandler{service: service}
}

// GetMastery handles GET /internal/learner/:user_id/mastery?skill=math_algebra
// @Summary Get user's mastery score for a skill
// @Description Returns the current mastery level for a user-skill combination
// @Tags learner
// @Produce json
// @Param user_id path string true "User ID"
// @Param skill query string true "Skill tag"
// @Success 200 {object} model.MasteryResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /internal/learner/{user_id}/mastery [get]
func (h *LearnerHandler) GetMastery(c *gin.Context) {
	userID := c.Param("user_id")
	skillTag := c.Query("skill")

	if userID == "" || skillTag == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user_id and skill are required",
		})
		return
	}

	mastery, err := h.service.GetMastery(userID, skillTag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, mastery)
}

// Health check endpoint
func (h *LearnerHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"service": "learner-model-service",
	})
}
