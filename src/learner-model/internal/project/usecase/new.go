package usecase

import (
	"time"

	"learner-model-service/internal/project"
	"learner-model-service/internal/project/repository"
	pkgLog "learner-model-service/pkg/log"
)

type usecase struct {
	l     pkgLog.Logger
	repo  repository.Repository
	clock func() time.Time
}

func New(l pkgLog.Logger, repo repository.Repository) project.UseCase {
	return &usecase{
		l:     l,
		repo:  repo,
		clock: time.Now,
	}
}
