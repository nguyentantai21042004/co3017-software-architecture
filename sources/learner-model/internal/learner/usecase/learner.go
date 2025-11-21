package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"learner-model-service/internal/learner"
	"learner-model-service/internal/learner/repository"
	"learner-model-service/internal/model"
)

const (
	MASTERY_THRESHOLD = 50
)

// GetMastery retrieves user's mastery score for a skill
func (uc *usecase) GetMastery(ctx context.Context, input learner.GetMasteryInput) (learner.MasteryOutput, error) {
	uc.l.Infof(ctx, "learner.usecase.GetMastery: starting | %s=%s | %s=%s",
		ErrCtxUserID, input.UserID, ErrCtxSkillTag, input.SkillTag)

	mastery, err := uc.repo.GetByUserAndSkill(ctx, input.UserID, input.SkillTag)
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		uc.l.Errorf(ctx, "learner.usecase.GetMastery: %s | %s=%s | %s=%s | error=%v",
			ErrMsgGetMasteryFailed, ErrCtxUserID, input.UserID, ErrCtxSkillTag, input.SkillTag, err)
		return learner.MasteryOutput{}, fmt.Errorf("%s: %w", ErrMsgGetMasteryFailed, err)
	}

	// Return default mastery if not found
	if mastery == nil {
		uc.l.Infof(ctx, "learner.usecase.GetMastery: mastery not found, returning default | %s=%s | %s=%s",
			ErrCtxUserID, input.UserID, ErrCtxSkillTag, input.SkillTag)
		return learner.MasteryOutput{
			UserID:       input.UserID,
			SkillTag:     input.SkillTag,
			MasteryScore: 0,
			LastUpdated:  time.Now().Format(time.RFC3339),
		}, nil
	}

	uc.l.Infof(ctx, "learner.usecase.GetMastery: success | %s=%s | %s=%s | %s=%d",
		ErrCtxUserID, input.UserID, ErrCtxSkillTag, input.SkillTag, ErrCtxScore, mastery.CurrentScore)

	return learner.MasteryOutput{
		UserID:       mastery.UserID,
		SkillTag:     mastery.SkillTag,
		MasteryScore: mastery.CurrentScore,
		LastUpdated:  mastery.LastUpdated.Format(time.RFC3339),
	}, nil
}

// UpdateMasteryFromEvent updates mastery based on submission event
func (uc *usecase) UpdateMasteryFromEvent(ctx context.Context, input learner.UpdateMasteryInput) error {
	uc.l.Infof(ctx, "learner.usecase.UpdateMasteryFromEvent: starting | %s=%s | %s=%s | score_obtained=%d",
		ErrCtxUserID, input.UserID, ErrCtxSkillTag, input.SkillTag, input.ScoreObtained)

	// Get current mastery
	currentMastery, err := uc.repo.GetByUserAndSkill(ctx, input.UserID, input.SkillTag)
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		uc.l.Errorf(ctx, "learner.usecase.UpdateMasteryFromEvent: %s | %s=%s | %s=%s | error=%v",
			ErrMsgUpdateMasteryFailed, ErrCtxUserID, input.UserID, ErrCtxSkillTag, input.SkillTag, err)
		return fmt.Errorf("%s: %w", ErrMsgUpdateMasteryFailed, err)
	}

	// Calculate new mastery score
	oldScore := 0
	if currentMastery != nil {
		oldScore = currentMastery.CurrentScore
	}

	// Formula: NewScore = (OldScore + ScoreObtained) / 2
	// This is a simple moving average
	newScore := (oldScore + input.ScoreObtained) / 2

	// Ensure score is within bounds [0, 100]
	if newScore < 0 {
		newScore = 0
	}
	if newScore > 100 {
		newScore = 100
	}

	uc.l.Infof(ctx, "learner.usecase.UpdateMasteryFromEvent: calculated new score | %s=%s | %s=%s | old=%d | obtained=%d | new=%d",
		ErrCtxUserID, input.UserID, ErrCtxSkillTag, input.SkillTag, oldScore, input.ScoreObtained, newScore)

	// Create or update mastery
	mastery := &model.SkillMastery{
		UserID:       input.UserID,
		SkillTag:     input.SkillTag,
		CurrentScore: newScore,
	}

	err = uc.repo.CreateOrUpdate(ctx, mastery)
	if err != nil {
		uc.l.Errorf(ctx, "learner.usecase.UpdateMasteryFromEvent: %s | %s=%s | %s=%s | %s=%d | error=%v",
			ErrMsgUpdateMasteryFailed, ErrCtxUserID, input.UserID, ErrCtxSkillTag, input.SkillTag, ErrCtxScore, newScore, err)
		return fmt.Errorf("%s: %w", ErrMsgUpdateMasteryFailed, err)
	}

	uc.l.Infof(ctx, "learner.usecase.UpdateMasteryFromEvent: success | %s=%s | %s=%s | %s=%d",
		ErrCtxUserID, input.UserID, ErrCtxSkillTag, input.SkillTag, ErrCtxScore, newScore)

	return nil
}
