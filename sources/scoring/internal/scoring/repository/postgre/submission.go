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
	"github.com/aarondl/sqlboiler/v4/queries/qm"
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

// FindAnsweredQuestionIDs retrieves a list of question IDs that a user has answered for a given skill tag.
func (r *submissionRepository) FindAnsweredQuestionIDs(ctx context.Context, userID, skillTag string) ([]int64, error) {
	var submissions sqlboiler.SubmissionSlice
	var err error

	// For simplicity, let's assume we want all questions submitted by the user.
	// A more robust implementation would need a mechanism to map questionID to skillTag.
	// Or, skillTag would need to be denormalized into the Submission model.
	// For now, skillTag parameter is ignored in the DB query.
	submissions, err = sqlboiler.Submissions(
		qm.Where("user_id = ?", userID),
		qm.Select("question_id"), // Only select the question_id
	).All(ctx, r.db)

	if err != nil {
		r.l.Errorf(ctx, "scoring.repository.postgre.FindAnsweredQuestionIDs: %s | user_id=%s | skill_tag=%s | error=%v",
			repository.ErrMsgDatabaseReadFailed, userID, skillTag, err)
		return nil, errors.Wrap(err, repository.ErrMsgDatabaseReadFailed)
	}

	// Extract unique question IDs
	questionIDs := make([]int64, 0, len(submissions))
	seen := make(map[int64]bool)
	for _, s := range submissions {
		if _, ok := seen[int64(s.QuestionID)]; !ok {
			seen[int64(s.QuestionID)] = true
			questionIDs = append(questionIDs, int64(s.QuestionID))
		}
	}

	r.l.Infof(ctx, "scoring.repository.postgre.FindAnsweredQuestionIDs: success | user_id=%s | skill_tag=%s | found_questions=%d",
		userID, skillTag, len(questionIDs))

	return questionIDs, nil
}