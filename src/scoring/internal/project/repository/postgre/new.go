package postgres

import (
	"database/sql"

	"scoring/internal/project/repository"
	"scoring/pkg/log"
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
