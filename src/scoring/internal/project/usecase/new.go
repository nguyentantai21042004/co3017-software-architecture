package usecase

import (
	"time"

	"scoring-serviceinternal/project"
	"scoring-serviceinternal/project/repository"
	pkgLog "scoring-servicepkg/log"
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
