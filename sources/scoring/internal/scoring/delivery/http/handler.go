package http

import (
	"net/http"
	"scoring/internal/scoring"
	"scoring/pkg/errors"
	"scoring/pkg/log"
	"scoring/pkg/response"


	"github.com/gin-gonic/gin"
)

type handler struct {
	l  log.Logger
	uc scoring.UseCase
}

// New creates a new HTTP handler for scoring module
func New(logger log.Logger, uc scoring.UseCase) Handler {
	return &handler{
		l:  logger,
		uc: uc,
	}
}

// Handler defines the HTTP handler interface for scoring module
type Handler interface {
	SubmitAnswer(c *gin.Context)
	Health(c *gin.Context)
	GetAnsweredQuestions(c *gin.Context) // Added for Content Service
}

// SubmitAnswer handles POST /api/scoring/submit
// @Summary Submit an answer for scoring
// @Description Submit a user's answer to a question and get immediate feedback
// @Tags scoring
// @Accept json
// @Produce json
// @Param request body SubmitRequest true "Submit Request"
// @Success 200 {object} response.Resp{data=SubmitResponse}
// @Failure 400 {object} response.Resp
// @Failure 500 {object} response.Resp
// @Router /scoring/submit [post]
func (h *handler) SubmitAnswer(c *gin.Context) {
	ctx := c.Request.Context()
	var req SubmitRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Errorf(ctx, "scoring.delivery.http.handler.SubmitAnswer: %s | error=%v | request_body=%+v",
			ErrMsgBindRequestFailed, err, req)
		httpErr := errors.NewHTTPError(http.StatusBadRequest, ErrMsgBindRequestFailed)
		httpErr.StatusCode = http.StatusBadRequest
		response.Error(c, httpErr, nil)
		return
	}

	// Validate required fields
	if req.UserID == "" {
		h.l.Errorf(ctx, "scoring.delivery.http.handler.SubmitAnswer: %s | user_id=empty", ErrMsgInvalidUserID)
		httpErr := errors.NewHTTPError(http.StatusBadRequest, ErrMsgInvalidUserID)
		httpErr.StatusCode = http.StatusBadRequest
		response.Error(c, httpErr, nil)
		return
	}

	if req.QuestionID <= 0 {
		h.l.Errorf(ctx, "scoring.delivery.http.handler.SubmitAnswer: %s | question_id=%d", ErrMsgInvalidQuestionID, req.QuestionID)
		httpErr := errors.NewHTTPError(http.StatusBadRequest, ErrMsgInvalidQuestionID)
		httpErr.StatusCode = http.StatusBadRequest
		response.Error(c, httpErr, nil)
		return
	}

	if req.Answer == "" {
		h.l.Errorf(ctx, "scoring.delivery.http.handler.SubmitAnswer: %s | answer=empty", ErrMsgInvalidAnswer)
		httpErr := errors.NewHTTPError(http.StatusBadRequest, ErrMsgInvalidAnswer)
		httpErr.StatusCode = http.StatusBadRequest
		response.Error(c, httpErr, nil)
		return
	}

	// Convert request to input
	input := req.ToSubmitInput()

	// Call use case
	output, err := h.uc.SubmitAnswer(ctx, input)
	if err != nil {
		h.l.Errorf(ctx, "scoring.delivery.http.handler.SubmitAnswer: %s | user_id=%s | question_id=%d | error=%v",
			ErrMsgSubmitAnswerFailed, req.UserID, req.QuestionID, err)
		httpErr := errors.NewHTTPError(http.StatusBadRequest, ErrMsgSubmitAnswerFailed)
		httpErr.StatusCode = http.StatusBadRequest
		response.Error(c, httpErr, nil)
		return
	}

	// Convert output to response
	resp := ToSubmitResponse(output)
	response.OK(c, resp)
}

// GetAnsweredQuestions handles GET /api/scoring/answered-questions
// @Summary Get answered questions by user and skill
// @Description Retrieves a list of question IDs that a user has answered for a given skill tag
// @Tags scoring
// @Accept json
// @Produce json
// @Param user_id query string true "User ID"
// @Param skill query string true "Skill Tag"
// @Success 200 {object} response.Resp{data=[]int64}
// @Failure 400 {object} response.Resp
// @Failure 500 {object} response.Resp
// @Router /scoring/answered-questions [get]
func (h *handler) GetAnsweredQuestions(c *gin.Context) {
	ctx := c.Request.Context()
	userID := c.Query("user_id")
	skillTag := c.Query("skill")

	if userID == "" {
		h.l.Errorf(ctx, "scoring.delivery.http.handler.GetAnsweredQuestions: %s | user_id=empty", ErrMsgInvalidUserID)
		httpErr := errors.NewHTTPError(http.StatusBadRequest, ErrMsgInvalidUserID)
		httpErr.StatusCode = http.StatusBadRequest
		response.Error(c, httpErr, nil)
		return
	}
	if skillTag == "" {
		h.l.Errorf(ctx, "scoring.delivery.http.handler.GetAnsweredQuestions: %s | skill=empty", ErrMsgInvalidSkillTag)
		httpErr := errors.NewHTTPError(http.StatusBadRequest, ErrMsgInvalidSkillTag)
		httpErr.StatusCode = http.StatusBadRequest
		response.Error(c, httpErr, nil)
		return
	}

	questionIDs, err := h.uc.GetAnsweredQuestions(ctx, userID, skillTag)
	if err != nil {
		h.l.Errorf(ctx, "scoring.delivery.http.handler.GetAnsweredQuestions: %s | user_id=%s | skill=%s | error=%v",
			ErrMsgGetAnsweredQuestionsFailed, userID, skillTag, err)
		// More specific error handling could be implemented based on error type from use case
		httpErr := errors.NewHTTPError(http.StatusInternalServerError, ErrMsgGetAnsweredQuestionsFailed)
		httpErr.StatusCode = http.StatusInternalServerError
		response.Error(c, httpErr, nil)
		return
	}

	response.OK(c, questionIDs)
}

// Health check endpoint
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
			"service": "scoring",
		},
	})
}