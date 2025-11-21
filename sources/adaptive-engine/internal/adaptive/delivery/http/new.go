package http

import (
	"adaptive-engine/internal/adaptive"
	"adaptive-engine/pkg/log"
)

type handler struct {
	l  log.Logger
	uc adaptive.UseCase
}

// New creates a new HTTP handler for adaptive module
func New(l log.Logger, uc adaptive.UseCase) Handler {
	return &handler{
		l:  l,
		uc: uc,
	}
}

// Handler defines the HTTP handler interface for adaptive module
type Handler interface {
	NextLesson(c any)
	Health(c any)
}
