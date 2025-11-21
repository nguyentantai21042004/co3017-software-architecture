package postgre

import (
	"context"
	"database/sql"
	"time"

	"scoring/internal/model"
	"scoring/internal/scoring/repository"
	"scoring/internal/sqlboiler"

	"scoring/pkg/log"

	"github.com/aarondl/null/v8"
	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/friendsofgo/errors"
)

type submissionRepository struct {
	db *sql.DB
	l  log.Logger
}

// New creates a new submission repository
func New(db *sql.DB, logger log.Logger) repository.Repository {
	return &submissionRepository{
		db: db,
		l:  logger,
	}
}

// Create inserts a new submission into the database using SQLBoiler
func (r *submissionRepository) Create(ctx context.Context, submission *model.Submission) error {
	// Convert domain model to SQLBoiler model
	boilerSubmission := &sqlboiler.Submission{
		UserID:          submission.UserID,
		QuestionID:      int(submission.QuestionID),
		SubmittedAnswer: submission.SubmittedAnswer,
		ScoreAwarded:    submission.ScoreAwarded,
		IsPassed:        submission.IsPassed,
		CreatedAt:       null.TimeFrom(time.Now()),
	}

	// Insert using SQLBoiler
	err := boilerSubmission.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		r.l.Errorf(ctx, "scoring.repository.postgre.Create: %s | user_id=%s | question_id=%d | error=%v",
			repository.ErrMsgDatabaseWriteFailed, submission.UserID, submission.QuestionID, err)
		return errors.Wrap(err, repository.ErrMsgDatabaseWriteFailed)
	}

	// Update domain model with generated ID and timestamp
	submission.ID = int64(boilerSubmission.ID)
	submission.CreatedAt = boilerSubmission.CreatedAt.Time

	r.l.Infof(ctx, "scoring.repository.postgre.Create: success | submission_id=%d | user_id=%s | question_id=%d | score=%d",
		submission.ID, submission.UserID, submission.QuestionID, submission.ScoreAwarded)

	return nil
}
