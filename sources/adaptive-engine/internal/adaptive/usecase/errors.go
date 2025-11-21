package usecase

// Error messages for usecase layer
const (
	ErrMsgFetchMasteryFailed      = "failed to fetch mastery from learner service"
	ErrMsgFetchContentFailed      = "failed to fetch content from content service"
	ErrMsgMasteryNotFound         = "mastery data not found for user"
	ErrMsgContentNotFound         = "content recommendation not found"
	ErrMsgInvalidMasteryScore     = "invalid mastery score received"
	ErrMsgInvalidContentResponse  = "invalid content response received"
	ErrMsgServiceTimeout          = "service request timeout"
	ErrMsgServiceConnectionFailed = "failed to connect to external service"
)

// Error context keys for detailed logging
const (
	ErrCtxUserID      = "user_id"
	ErrCtxSkillTag    = "skill_tag"
	ErrCtxContentType = "content_type"
	ErrCtxStatusCode  = "status_code"
	ErrCtxServiceURL  = "service_url"
	ErrCtxErrorDetail = "error_detail"
)
