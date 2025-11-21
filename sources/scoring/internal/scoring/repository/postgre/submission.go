package postgre

import (
	"database/sql"
	"scoring/internal/model"
	"scoring/internal/scoring/repository"
	"time"
)

type submissionRepository struct {
	db *sql.DB
}

// New creates a new submission repository
func New(db *sql.DB) repository.Repository {
	return &submissionRepository{db: db}
}

// Create inserts a new submission into the database
func (r *submissionRepository) Create(submission *model.Submission) error {
	query := `
		INSERT INTO submissions (user_id, question_id, submitted_answer, score_awarded, is_passed, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	submission.CreatedAt = time.Now()

	err := r.db.QueryRow(
		query,
		submission.UserID,
		submission.QuestionID,
		submission.SubmittedAnswer,
		submission.ScoreAwarded,
		submission.IsPassed,
		submission.CreatedAt,
	).Scan(&submission.ID)

	if err != nil {
		return err
	}

	return nil
}
