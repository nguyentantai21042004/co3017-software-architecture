package repository

import (
	"context"
	"scoring/internal/model"
)

// Repository defines the interface for submission repository
type Repository interface {
	Create(ctx context.Context, submission *model.Submission) error
	// Add more methods as needed
}
