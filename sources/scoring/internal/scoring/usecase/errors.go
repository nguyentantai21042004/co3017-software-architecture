package usecase

// Error messages for usecase layer
const (
	ErrMsgFetchQuestionFailed     = "failed to fetch question from content service"
	ErrMsgSaveSubmissionFailed    = "failed to save submission to database"
	ErrMsgPublishEventFailed      = "failed to publish submission event"
	ErrMsgInvalidAnswer           = "invalid answer format"
	ErrMsgQuestionNotFound        = "question not found"
	ErrMsgServiceTimeout          = "service request timeout"
	ErrMsgServiceConnectionFailed = "failed to connect to external service"
	ErrMsgGetAnsweredQuestionsFailed = "failed to retrieve answered questions from database"
)

// Error context keys for detailed logging
const (
	ErrCtxUserID       = "user_id"
	ErrCtxQuestionID   = "question_id"
	ErrCtxAnswer       = "answer"
	ErrCtxScore        = "score"
	ErrCtxSkillTag     = "skill_tag"
	ErrCtxIsCorrect    = "is_correct"
	ErrCtxErrorDetail  = "error_detail"
	ErrCtxSubmissionID = "submission_id"
)
