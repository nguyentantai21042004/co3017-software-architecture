package adaptive

import "errors"

var (
	ErrInvalidRequest     = errors.New("invalid request")
	ErrServiceUnavailable = errors.New("service unavailable")
	ErrMasteryThreshold   = errors.New("mastery score below threshold")
	ErrContentNotFound    = errors.New("content not found")
	ErrContentUnavailable = errors.New("content unavailable")
	ErrContentInvalid     = errors.New("content invalid")
	ErrContentTimeout     = errors.New("content timeout")
	ErrContentCanceled    = errors.New("content canceled")
	ErrContentAborted     = errors.New("content aborted")
	ErrContentFailed      = errors.New("content failed")
)
