package http

// Error messages for HTTP handler layer
const (
	ErrMsgBindRequestFailed       = "failed to bind request body: invalid JSON format or missing required fields"
	ErrMsgRecommendNextLesson     = "failed to recommend next lesson: internal service error"
	ErrMsgInvalidUserID           = "invalid user_id: must be non-empty string"
	ErrMsgInvalidCurrentSkill     = "invalid current_skill: must be non-empty string"
	ErrMsgRequestValidationFailed = "request validation failed: check request body format"
)
