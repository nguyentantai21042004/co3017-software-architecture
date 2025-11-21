package curl

import "errors"

var (
	// ErrServiceUnavailable is returned when the external service is unavailable
	ErrServiceUnavailable = errors.New("service unavailable")

	// ErrInvalidResponse is returned when the response cannot be parsed
	ErrInvalidResponse = errors.New("invalid response from service")

	// ErrNotFound is returned when the resource is not found
	ErrNotFound = errors.New("resource not found")

	// ErrBadRequest is returned when the request is invalid
	ErrBadRequest = errors.New("bad request")
)
