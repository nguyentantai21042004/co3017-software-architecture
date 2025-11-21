package usecase

import (
	"scoring/internal/publisher"
	"scoring/internal/scoring"
	"scoring/internal/scoring/repository"
	"scoring/pkg/curl"
	"scoring/pkg/log"
)

type usecase struct {
	l             log.Logger
	repo          repository.Repository
	publisher     publisher.EventPublisher
	contentClient curl.ContentServiceClientInterface
}

// New creates a new scoring usecase
func New(
	logger log.Logger,
	repo repository.Repository,
	pub publisher.EventPublisher,
	contentClient curl.ContentServiceClientInterface,
) scoring.UseCase {
	return &usecase{
		l:             logger,
		repo:          repo,
		publisher:     pub,
		contentClient: contentClient,
	}
}
