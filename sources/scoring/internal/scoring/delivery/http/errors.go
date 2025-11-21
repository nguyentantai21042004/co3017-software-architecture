package http

// Error messages for HTTP handler layer
const (
	ErrMsgBindRequestFailed       = "failed to bind request body: invalid JSON format or missing required fields"
	ErrMsgSubmitAnswerFailed      = "failed to submit answer: internal service error"
	ErrMsgInvalidUserID           = "invalid user_id: must be non-empty string"
	ErrMsgInvalidQuestionID       = "invalid question_id: must be positive integer"
	ErrMsgInvalidAnswer           = "invalid answer: must be non-empty string"
	ErrMsgRequestValidationFailed = "request validation failed: check request body format"
)
