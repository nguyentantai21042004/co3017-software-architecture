package postgre

import (
	"context"
	"database/sql"
	"time"

	"learner-model-service/internal/learner/repository"
	"learner-model-service/internal/model"
	"learner-model-service/internal/sqlboiler"

	"github.com/aarondl/null/v8"
	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/friendsofgo/errors"
)

// GetByUserAndSkill retrieves mastery for a user and skill using SQLBoiler
func (r *implRepository) GetByUserAndSkill(ctx context.Context, userID, skillTag string) (*model.SkillMastery, error) {
	// Use SQLBoiler to find the record
	boilerMastery, err := sqlboiler.FindSkillMastery(ctx, r.db, userID, skillTag)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		r.l.Errorf(ctx, "learner.repository.postgre.GetByUserAndSkill: %s | user_id=%s | skill_tag=%s | error=%v",
			repository.ErrMsgDatabaseQueryFailed, userID, skillTag, err)
		return nil, repository.ErrDatabaseFailure
	}

	// Convert SQLBoiler model to domain model
	domainMastery := &model.SkillMastery{
		UserID:       boilerMastery.UserID,
		SkillTag:     boilerMastery.SkillTag,
		CurrentScore: int(boilerMastery.CurrentScore.Int), // Convert null.Int to int
		LastUpdated:  boilerMastery.LastUpdated.Time,      // Convert null.Time to time.Time
	}

	r.l.Infof(ctx, "learner.repository.postgre.GetByUserAndSkill: success | user_id=%s | skill_tag=%s | score=%d",
		userID, skillTag, domainMastery.CurrentScore)

	return domainMastery, nil
}

// CreateOrUpdate creates or updates mastery record using SQLBoiler
func (r *implRepository) CreateOrUpdate(ctx context.Context, mastery *model.SkillMastery) error {
	// Convert domain model to SQLBoiler model
	boilerMastery := &sqlboiler.SkillMastery{
		UserID:       mastery.UserID,
		SkillTag:     mastery.SkillTag,
		CurrentScore: null.IntFrom(mastery.CurrentScore), // Convert int to null.Int
		LastUpdated:  null.TimeFrom(time.Now()),          // Set current time
	}

	// Use SQLBoiler's Upsert (INSERT ... ON CONFLICT ... DO UPDATE)
	err := boilerMastery.Upsert(
		ctx,
		r.db,
		true,         // updateOnConflict
		nil,          // conflictColumns (nil means use primary key)
		boil.Infer(), // updateColumns
		boil.Infer(), // insertColumns
	)

	if err != nil {
		r.l.Errorf(ctx, "learner.repository.postgre.CreateOrUpdate: %s | user_id=%s | skill_tag=%s | score=%d | error=%v",
			repository.ErrMsgDatabaseWriteFailed, mastery.UserID, mastery.SkillTag, mastery.CurrentScore, err)
		return repository.ErrDatabaseFailure
	}

	r.l.Infof(ctx, "learner.repository.postgre.CreateOrUpdate: success | user_id=%s | skill_tag=%s | score=%d",
		mastery.UserID, mastery.SkillTag, mastery.CurrentScore)

	// Update the domain model with the timestamp from DB
	mastery.LastUpdated = boilerMastery.LastUpdated.Time

	return nil
}
