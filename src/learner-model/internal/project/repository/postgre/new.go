package postgres

import (
	"database/sql"

	"learner-model-service/internal/project/repository"
	"learner-model-service/pkg/log"
)

type implRepository struct {
	db *sql.DB
	l  log.Logger
}

// New creates a new PostgreSQL repository for projects
func New(db *sql.DB, l log.Logger) repository.Repository {
	return &implRepository{
		db: db,
		l:  l,
	}
}
