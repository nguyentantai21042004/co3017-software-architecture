package http

import (
	"net/http"

	"learner-model-service/pkg/response"

	"github.com/gin-gonic/gin"
)

// GetMastery handles GET /internal/learner/:user_id/mastery?skill=skill_tag
func (h *handler) GetMastery(c any) {
	ctx := c.(*gin.Context)

	userID := ctx.Param("user_id")
	skillTag := ctx.Query("skill")

	if userID == "" || skillTag == "" {
		response.Error(ctx, http.StatusBadRequest, ErrMsgInvalidRequest, nil)
		return
	}

	// Create request and convert to input
	req := GetMasteryRequest{
		UserID:   userID,
		SkillTag: skillTag,
	}
	input := req.ToGetMasteryInput()

	// Call use case
	output, err := h.uc.GetMastery(ctx.Request.Context(), input)
	if err != nil {
		h.l.Errorf(ctx.Request.Context(), "learner.handler.GetMastery: failed | error=%v", err)
		response.Error(ctx, http.StatusInternalServerError, ErrMsgInternalError, nil)
		return
	}

	// Convert output to response
	resp := ToMasteryResponse(output)
	response.OK(ctx, resp)
}

// Health handles GET /health
func (h *handler) Health(c any) {
	ctx := c.(*gin.Context)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"service": "learner-model-service",
	})
}
