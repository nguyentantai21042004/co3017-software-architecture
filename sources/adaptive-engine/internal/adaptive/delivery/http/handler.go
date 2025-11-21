package http

import (
	"net/http"

	"adaptive-engine/pkg/errors"
	"adaptive-engine/pkg/response"

	"github.com/gin-gonic/gin"
)

// @Summary      Recommend next lesson
// @Description  Returns recommended next lesson for a user based on current knowledge state.
// @Tags         Adaptive
// @Accept       json
// @Produce      json
// @Param        request  body      NextLessonRequest  true  "Recommendation request"
// @Success      200      {object}  response.Resp{data=NextLessonResponse}
// @Failure      400      {object}  response.Resp
// @Failure      500      {object}  response.Resp
// @Router       /next-lesson [post]
func (h *handler) NextLesson(c any) {
	ctx := c.(*gin.Context)
	reqCtx := ctx.Request.Context()

	var req NextLessonRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.l.Errorf(reqCtx, "adaptive.delivery.http.handler.NextLesson: %s | error=%v | request_body=%+v",
			ErrMsgBindRequestFailed, err, req)
		httpErr := errors.NewHTTPError(http.StatusBadRequest, ErrMsgBindRequestFailed)
		httpErr.StatusCode = http.StatusBadRequest
		response.Error(ctx, httpErr, nil)
		return
	}

	// Validate required fields
	if req.UserID == "" {
		h.l.Errorf(reqCtx, "adaptive.delivery.http.handler.NextLesson: %s | user_id=empty", ErrMsgInvalidUserID)
		httpErr := errors.NewHTTPError(http.StatusBadRequest, ErrMsgInvalidUserID)
		httpErr.StatusCode = http.StatusBadRequest
		response.Error(ctx, httpErr, nil)
		return
	}

	if req.CurrentSkill == "" {
		h.l.Errorf(reqCtx, "adaptive.delivery.http.handler.NextLesson: %s | current_skill=empty", ErrMsgInvalidCurrentSkill)
		httpErr := errors.NewHTTPError(http.StatusBadRequest, ErrMsgInvalidCurrentSkill)
		httpErr.StatusCode = http.StatusBadRequest
		response.Error(ctx, httpErr, nil)
		return
	}

	o, err := h.uc.RecommendNextLesson(reqCtx, req.ToRecommendInput())
	if err != nil {
		h.l.Errorf(reqCtx, "adaptive.delivery.http.handler.NextLesson: %s | user_id=%s | skill=%s | error=%v",
			ErrMsgRecommendNextLesson, req.UserID, req.CurrentSkill, err)
		httpErr := errors.NewHTTPError(http.StatusBadRequest, ErrMsgRecommendNextLesson)
		httpErr.StatusCode = http.StatusBadRequest
		response.Error(ctx, httpErr, nil)
		return
	}

	response.OK(ctx, toResponse(o))
}

// @Summary      Health check
// @Description  Show the service health and status
// @Tags         Health
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /health [get]
func (h *handler) Health(c any) {
	ctx := c.(*gin.Context)
	ctx.JSON(http.StatusOK, response.Resp{
		ErrorCode: 0,
		Message:   "Healthy",
		Data: map[string]interface{}{
			"status":  "healthy",
			"service": "adaptive-engine",
		},
	})
}
