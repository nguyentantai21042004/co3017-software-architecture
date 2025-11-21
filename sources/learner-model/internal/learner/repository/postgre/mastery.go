package postgre

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"learner-model-service/internal/learner/repository"
	"learner-model-service/internal/model"
)

// GetByUserAndSkill retrieves mastery for a user and skill
func (r *implRepository) GetByUserAndSkill(ctx context.Context, userID, skillTag string) (*model.SkillMastery, error) {
	query := `
		SELECT user_id, skill_tag, current_score, last_updated
		FROM skill_mastery
		WHERE user_id = $1 AND skill_tag = $2
	`

	var mastery model.SkillMastery
	err := r.db.QueryRowContext(ctx, query, userID, skillTag).Scan(
		&mastery.UserID,
		&mastery.SkillTag,
		&mastery.CurrentScore,
		&mastery.LastUpdated,
	)

	if err == sql.ErrNoRows {
		return nil, repository.ErrNotFound
	}

	if err != nil {
		r.l.Errorf(ctx, "learner.repository.GetByUserAndSkill: failed | user_id=%s | skill_tag=%s | error=%v",
			userID, skillTag, err)
		return nil, fmt.Errorf("%w: %v", repository.ErrDatabaseFailure, err)
	}

	return &mastery, nil
}

// CreateOrUpdate creates or updates mastery record
func (r *implRepository) CreateOrUpdate(ctx context.Context, mastery *model.SkillMastery) error {
	query := `
		INSERT INTO skill_mastery (user_id, skill_tag, current_score, last_updated)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (user_id, skill_tag)
		DO UPDATE SET
			current_score = EXCLUDED.current_score,
			last_updated = EXCLUDED.last_updated
	`

	mastery.LastUpdated = time.Now()

	_, err := r.db.ExecContext(
		ctx,
		query,
		mastery.UserID,
		mastery.SkillTag,
		mastery.CurrentScore,
		mastery.LastUpdated,
	)

	if err != nil {
		r.l.Errorf(ctx, "learner.repository.CreateOrUpdate: failed | user_id=%s | skill_tag=%s | score=%d | error=%v",
			mastery.UserID, mastery.SkillTag, mastery.CurrentScore, err)
		return fmt.Errorf("%w: %v", repository.ErrDatabaseFailure, err)
	}

	return nil
}
