package scoring

import "errors"

var (
	ErrInvalidRequest       = errors.New("invalid request")
	ErrQuestionNotFound     = errors.New("question not found")
	ErrServiceUnavailable   = errors.New("service unavailable")
	ErrSubmissionSaveFailed = errors.New("failed to save submission")
	ErrEventPublishFailed   = errors.New("failed to publish event")
)
