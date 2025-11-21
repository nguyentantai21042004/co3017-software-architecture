package learner

import "errors"

var (
	ErrMasteryNotFound = errors.New("mastery not found")
	ErrInvalidInput    = errors.New("invalid input")
	ErrUpdateFailed    = errors.New("failed to update mastery")
)
