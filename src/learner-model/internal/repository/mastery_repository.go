package repository

import (
	"database/sql"
	"learner-model-service/internal/model"
	"time"
)

type MasteryRepository interface {
	GetByUserAndSkill(userID, skillTag string) (*model.SkillMastery, error)
	UpdateScore(userID, skillTag string, newScore int) error
	CreateOrUpdate(mastery *model.SkillMastery) error
}

type masteryRepository struct {
	db *sql.DB
}

func NewMasteryRepository(db *sql.DB) MasteryRepository {
	return &masteryRepository{db: db}
}

func (r *masteryRepository) GetByUserAndSkill(userID, skillTag string) (*model.SkillMastery, error) {
	query := `
		SELECT user_id, skill_tag, current_score, last_updated
		FROM skill_mastery
		WHERE user_id = $1 AND skill_tag = $2
	`

	var mastery model.SkillMastery
	err := r.db.QueryRow(query, userID, skillTag).Scan(
		&mastery.UserID,
		&mastery.SkillTag,
		&mastery.CurrentScore,
		&mastery.LastUpdated,
	)

	if err == sql.ErrNoRows {
		return nil, nil // Not found
	}

	if err != nil {
		return nil, err
	}

	return &mastery, nil
}

func (r *masteryRepository) UpdateScore(userID, skillTag string, newScore int) error {
	query := `
		UPDATE skill_mastery
		SET current_score = $1, last_updated = $2
		WHERE user_id = $3 AND skill_tag = $4
	`

	_, err := r.db.Exec(query, newScore, time.Now(), userID, skillTag)
	return err
}

func (r *masteryRepository) CreateOrUpdate(mastery *model.SkillMastery) error {
	query := `
		INSERT INTO skill_mastery (user_id, skill_tag, current_score, last_updated)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (user_id, skill_tag)
		DO UPDATE SET
			current_score = EXCLUDED.current_score,
			last_updated = EXCLUDED.last_updated
	`

	mastery.LastUpdated = time.Now()

	_, err := r.db.Exec(
		query,
		mastery.UserID,
		mastery.SkillTag,
		mastery.CurrentScore,
		mastery.LastUpdated,
	)

	return err
}
