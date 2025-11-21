package usecase

import (
	"learner-model-service/internal/learner"
	"learner-model-service/internal/learner/repository"
	"learner-model-service/pkg/log"
)

type usecase struct {
	l    log.Logger
	repo repository.Repository
}

// New creates a new learner usecase
func New(l log.Logger, repo repository.Repository) learner.UseCase {
	return &usecase{
		l:    l,
		repo: repo,
	}
}
