package repository

import "scoring/internal/model"

// Repository defines the interface for submission repository
type Repository interface {
	Create(submission *model.Submission) error
	// Add more methods as needed
}
