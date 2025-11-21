package http

import (
	"learner-model-service/internal/learner"
	"learner-model-service/pkg/log"

	"github.com/gin-gonic/gin"
)

type handler struct {
	l  log.Logger
	uc learner.UseCase
}

// New creates a new HTTP handler for learner module
func New(l log.Logger, uc learner.UseCase) Handler {
	return &handler{
		l:  l,
		uc: uc,
	}
}

// Handler defines the HTTP handler interface for learner module
type Handler interface {
	GetMastery(c *gin.Context)
	UpdateMasteryFromEvent(c *gin.Context)
	Health(c *gin.Context)
}
