package service

import (
	"fmt"
	"log"
	"time"

	"learner-model-service/internal/model"
	"learner-model-service/internal/repository"
)

type LearnerService interface {
	GetMastery(userID, skillTag string) (*model.MasteryResponse, error)
	UpdateMasteryFromEvent(event *model.SubmissionEvent) error
}

type learnerService struct {
	repo repository.MasteryRepository
}

func NewLearnerService(repo repository.MasteryRepository) LearnerService {
	return &learnerService{repo: repo}
}

func (s *learnerService) GetMastery(userID, skillTag string) (*model.MasteryResponse, error) {
	mastery, err := s.repo.GetByUserAndSkill(userID, skillTag)
	if err != nil {
		return nil, fmt.Errorf("failed to get mastery: %w", err)
	}

	if mastery == nil {
		// Return default mastery if not found
		return &model.MasteryResponse{
			UserID:       userID,
			SkillTag:     skillTag,
			MasteryScore: 0,
			LastUpdated:  time.Now().Format(time.RFC3339),
		}, nil
	}

	return &model.MasteryResponse{
		UserID:       mastery.UserID,
		SkillTag:     mastery.SkillTag,
		MasteryScore: mastery.CurrentScore,
		LastUpdated:  mastery.LastUpdated.Format(time.RFC3339),
	}, nil
}

func (s *learnerService) UpdateMasteryFromEvent(event *model.SubmissionEvent) error {
	log.Printf("🧮 Calculating new mastery for user: %s, skill: %s, score_obtained: %d",
		event.UserID, event.SkillTag, event.ScoreObtained)

	// Get current mastery
	currentMastery, err := s.repo.GetByUserAndSkill(event.UserID, event.SkillTag)
	if err != nil {
		return fmt.Errorf("failed to get current mastery: %w", err)
	}

	// Calculate new mastery score
	oldScore := 0
	if currentMastery != nil {
		oldScore = currentMastery.CurrentScore
	}

	// Formula: NewScore = (OldScore + ScoreObtained) / 2
	// This is a simple moving average
	newScore := (oldScore + event.ScoreObtained) / 2

	// Ensure score is within bounds [0, 100]
	if newScore < 0 {
		newScore = 0
	}
	if newScore > 100 {
		newScore = 100
	}

	log.Printf("📊 Mastery update: %s [%s] - Old: %d, Obtained: %d, New: %d",
		event.UserID, event.SkillTag, oldScore, event.ScoreObtained, newScore)

	// Create or update mastery
	mastery := &model.SkillMastery{
		UserID:       event.UserID,
		SkillTag:     event.SkillTag,
		CurrentScore: newScore,
	}

	err = s.repo.CreateOrUpdate(mastery)
	if err != nil {
		return fmt.Errorf("failed to update mastery: %w", err)
	}

	log.Printf("✅ Updated mastery: user=%s, skill=%s, new_score=%d",
		event.UserID, event.SkillTag, newScore)

	return nil
}
