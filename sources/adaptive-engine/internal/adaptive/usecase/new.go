package usecase

import (
	"adaptive-engine/internal/adaptive"
	"adaptive-engine/pkg/curl"
	"adaptive-engine/pkg/log"
)

type usecase struct {
	l             log.Logger
	learnerClient curl.LearnerServiceClientInterface
	contentClient curl.ContentServiceClientInterface
}

// New creates a new adaptive usecase
func New(l log.Logger, learnerClient curl.LearnerServiceClientInterface, contentClient curl.ContentServiceClientInterface) adaptive.UseCase {
	return &usecase{
		l:             l,
		learnerClient: learnerClient,
		contentClient: contentClient,
	}
}
