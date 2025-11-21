package postgre

import (
	"database/sql"
	"learner-model-service/internal/learner/repository"
	"learner-model-service/pkg/log"
)

type implRepository struct {
	db *sql.DB
	l  log.Logger
}

// New creates a new PostgreSQL repository
func New(db *sql.DB, l log.Logger) repository.Repository {
	return &implRepository{
		db: db,
		l:  l,
	}
}
