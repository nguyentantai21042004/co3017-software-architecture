package repository

import (
	"context"
	"scoring/internal/model"
)

// Repository defines the interface for submission repository
type Repository interface {
	Create(ctx context.Context, submission *model.Submission) error
	FindAnsweredQuestionIDs(ctx context.Context, userID, skillTag string) ([]int64, error)
	// Add more methods as needed
}
