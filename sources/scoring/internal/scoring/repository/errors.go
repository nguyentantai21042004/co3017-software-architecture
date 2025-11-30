package repository

import "errors"

var (
	ErrNotFound        = errors.New("submission not found")
	ErrDatabaseFailure = errors.New("database operation failed")
)

// Error messages for detailed logging
const (
	ErrMsgDatabaseQueryFailed = "database query failed"
	ErrMsgDatabaseWriteFailed = "database write operation failed"
	ErrMsgDatabaseReadFailed  = "database read operation failed"
	ErrMsgRecordNotFound      = "record not found in database"
)
