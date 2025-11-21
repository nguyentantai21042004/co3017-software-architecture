package usecase

import (
	"time"

	"adaptive-engine-service/internal/project"
	"adaptive-engine-service/internal/project/repository"
	pkgLog "adaptive-engine/pkg/log"
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
