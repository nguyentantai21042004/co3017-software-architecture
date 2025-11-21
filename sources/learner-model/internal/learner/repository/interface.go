package repository

import (
	"context"
	"learner-model-service/internal/model"
)

//go:generate mockery --name Repository
type Repository interface {
	GetByUserAndSkill(ctx context.Context, userID, skillTag string) (*model.SkillMastery, error)
	CreateOrUpdate(ctx context.Context, mastery *model.SkillMastery) error
}
