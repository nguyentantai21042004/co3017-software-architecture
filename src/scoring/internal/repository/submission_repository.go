package repository

import (
	"database/sql"
	"scoring-serviceinternal/model"
	"time"
)

type SubmissionRepository interface {
	Create(submission *model.Submission) error
}

type submissionRepository struct {
	db *sql.DB
}

func NewSubmissionRepository(db *sql.DB) SubmissionRepository {
	return &submissionRepository{db: db}
}

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

	return err
}
