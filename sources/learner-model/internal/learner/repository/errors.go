package repository

import "errors"

var (
	ErrNotFound        = errors.New("mastery not found")
	ErrDatabaseFailure = errors.New("database operation failed")
)
