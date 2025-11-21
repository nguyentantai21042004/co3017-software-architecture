package handler

import (
	"net/http"
	"scoring-serviceinternal/model"
	"scoring-serviceinternal/service"

	"github.com/gin-gonic/gin"
)

type ScoringHandler struct {
	service service.ScoringService
}

func NewScoringHandler(service service.ScoringService) *ScoringHandler {
	return &ScoringHandler{service: service}
}

// SubmitAnswer handles POST /api/scoring-servicesubmit
// @Summary Submit an answer for scoring
// @Description Submit a user's answer to a question and get immediate feedback
// @Tags scoring
// @Accept json
// @Produce json
// @Param request body model.SubmitRequest true "Submit Request"
// @Success 200 {object} model.SubmitResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/scoring-servicesubmit [post]
func (h *ScoringHandler) SubmitAnswer(c *gin.Context) {
	var req model.SubmitRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request payload",
		})
		return
	}

	response, err := h.service.SubmitAnswer(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// Health check endpoint
func (h *ScoringHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"service": "scoring",
	})
}
