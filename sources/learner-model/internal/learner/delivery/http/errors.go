package http

// Error messages for HTTP handler layer
const (
	ErrMsgBindRequestFailed   = "failed to bind request body: invalid JSON format or missing required fields"
	ErrMsgGetMasteryFailed    = "failed to get mastery level: internal service error"
	ErrMsgUpdateMasteryFailed = "failed to update mastery level: internal service error"
	ErrMsgInvalidUserID       = "invalid user_id: must be non-empty string"
	ErrMsgInvalidSkillTag     = "invalid skill_tag: must be non-empty string"
	ErrMsgInternalError       = "internal server error"
)
