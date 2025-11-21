package repository

import "errors"

var (
	// ErrNotFound is returned when a record is not found
	ErrNotFound = errors.New("record not found")

	// ErrDuplicateEntry is returned when trying to create a duplicate entry
	ErrDuplicateEntry = errors.New("duplicate entry")

	// ErrDatabaseConnection is returned when database connection fails
	ErrDatabaseConnection = errors.New("database connection failed")
)
