package http

import (
	"net/http"

	"learner-model-service/pkg/errors"
	"learner-model-service/pkg/response"

	"github.com/gin-gonic/gin"
)

// GetMastery handles GET /internal/learner/:user_id/mastery?skill=skill_tag
// @Summary Get mastery level
// @Description Get the mastery level for a specific skill of a user
// @Tags learner
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Param skill query string true "Skill Tag"
// @Success 200 {object} response.Resp{data=MasteryResponse}
// @Failure 400 {object} response.Resp
// @Failure 500 {object} response.Resp
// @Router /learner/{user_id}/mastery [get]
func (h *handler) GetMastery(c *gin.Context) {
	ctx := c.Request.Context()

	userID := c.Param("user_id")
	skillTag := c.Query("skill")

	// Validate required parameters
	if userID == "" {
		h.l.Errorf(ctx, "learner.delivery.http.handler.GetMastery: %s | user_id=empty", ErrMsgInvalidUserID)
		httpErr := errors.NewHTTPError(http.StatusBadRequest, ErrMsgInvalidUserID)
		httpErr.StatusCode = http.StatusBadRequest
		response.Error(c, httpErr, nil)
		return
	}

	if skillTag == "" {
		h.l.Errorf(ctx, "learner.delivery.http.handler.GetMastery: %s | skill_tag=empty", ErrMsgInvalidSkillTag)
		httpErr := errors.NewHTTPError(http.StatusBadRequest, ErrMsgInvalidSkillTag)
		httpErr.StatusCode = http.StatusBadRequest
		response.Error(c, httpErr, nil)
		return
	}

	// Create request and convert to input
	req := GetMasteryRequest{
		UserID:   userID,
		SkillTag: skillTag,
	}
	input := req.ToGetMasteryInput()

	// Call use case
	output, err := h.uc.GetMastery(ctx, input)
	if err != nil {
		h.l.Errorf(ctx, "learner.delivery.http.handler.GetMastery: %s | user_id=%s | skill_tag=%s | error=%v",
			ErrMsgGetMasteryFailed, userID, skillTag, err)
		httpErr := errors.NewHTTPError(http.StatusInternalServerError, ErrMsgGetMasteryFailed)
		httpErr.StatusCode = http.StatusInternalServerError
		response.Error(c, httpErr, nil)
		return
	}

	// Convert output to response
	resp := ToMasteryResponse(output)
	response.OK(c, resp)
}

// UpdateMasteryFromEvent handles internal mastery updates from RabbitMQ events
func (h *handler) UpdateMasteryFromEvent(c *gin.Context) {
	ctx := c.Request.Context()
	var req UpdateMasteryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Errorf(ctx, "learner.delivery.http.handler.UpdateMasteryFromEvent: %s | error=%v | request_body=%+v",
			ErrMsgBindRequestFailed, err, req)
		httpErr := errors.NewHTTPError(http.StatusBadRequest, ErrMsgBindRequestFailed)
		httpErr.StatusCode = http.StatusBadRequest
		response.Error(c, httpErr, nil)
		return
	}

	// Validate required fields
	if req.UserID == "" {
		h.l.Errorf(ctx, "learner.delivery.http.handler.UpdateMasteryFromEvent: %s | user_id=empty", ErrMsgInvalidUserID)
		httpErr := errors.NewHTTPError(http.StatusBadRequest, ErrMsgInvalidUserID)
		httpErr.StatusCode = http.StatusBadRequest
		response.Error(c, httpErr, nil)
		return
	}

	if req.SkillTag == "" {
		h.l.Errorf(ctx, "learner.delivery.http.handler.UpdateMasteryFromEvent: %s | skill_tag=empty", ErrMsgInvalidSkillTag)
		httpErr := errors.NewHTTPError(http.StatusBadRequest, ErrMsgInvalidSkillTag)
		httpErr.StatusCode = http.StatusBadRequest
		response.Error(c, httpErr, nil)
		return
	}

	// Convert to input
	input := req.ToUpdateMasteryInput()

	// Call use case
	err := h.uc.UpdateMasteryFromEvent(ctx, input)
	if err != nil {
		h.l.Errorf(ctx, "learner.delivery.http.handler.UpdateMasteryFromEvent: %s | user_id=%s | skill_tag=%s | error=%v",
			ErrMsgUpdateMasteryFailed, req.UserID, req.SkillTag, err)
		httpErr := errors.NewHTTPError(http.StatusInternalServerError, ErrMsgUpdateMasteryFailed)
		httpErr.StatusCode = http.StatusInternalServerError
		response.Error(c, httpErr, nil)
		return
	}

	response.OK(c, map[string]interface{}{
		"message": "Mastery updated successfully",
	})
}

// Health handles GET /health
// @Summary Health check
// @Description Show the service health and status
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} response.Resp
// @Router /health [get]
func (h *handler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, response.Resp{
		ErrorCode: 0,
		Message:   "Healthy",
		Data: map[string]interface{}{
			"status":  "healthy",
			"service": "learner-model",
		},
	})
}
