package service

import (
	"fmt"
	"log"
	"scoring/internal/model"
	"scoring/internal/publisher"
	"scoring/internal/repository"
	"time"
)

type ScoringService interface {
	SubmitAnswer(req *model.SubmitRequest) (*model.SubmitResponse, error)
}

type scoringService struct {
	repo          repository.SubmissionRepository
	publisher     publisher.EventPublisher
	contentClient ContentClient
}

func NewScoringService(repo repository.SubmissionRepository, pub publisher.EventPublisher, client ContentClient) ScoringService {
	return &scoringService{
		repo:          repo,
		publisher:     pub,
		contentClient: client,
	}
}

// Question struct to parse Content Service response
type QuestionData struct {
	ID            int64  `json:"id"`
	Content       string `json:"content"`
	CorrectAnswer string `json:"correct_answer"`
	SkillTag      string `json:"skill_tag"`
	IsRemedial    bool   `json:"is_remedial"`
}

type ContentServiceResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    QuestionData `json:"data"`
}

func (s *scoringService) SubmitAnswer(req *model.SubmitRequest) (*model.SubmitResponse, error) {
	log.Printf("📝 Processing submission: user=%s, question=%d, answer=%s",
		req.UserID, req.QuestionID, req.Answer)

	// Step 1: Fetch correct answer from Content Service
	correctAnswer, skillTag, err := s.contentClient.FetchQuestion(req.QuestionID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch question: %w", err)
	}

	// Step 2: Score the answer
	isCorrect := req.Answer == correctAnswer
	score := 0
	if isCorrect {
		score = 100
	}
	isPassed := score >= 50

	feedback := "Sai rồi, hãy thử lại!"
	if isCorrect {
		feedback = "Chính xác! Bạn đã trả lời đúng."
	}

	log.Printf("🎯 Scoring result: correct=%v, score=%d", isCorrect, score)

	// Step 3: Save to database
	submission := &model.Submission{
		UserID:          req.UserID,
		QuestionID:      req.QuestionID,
		SubmittedAnswer: req.Answer,
		ScoreAwarded:    score,
		IsPassed:        isPassed,
	}

	err = s.repo.Create(submission)
	if err != nil {
		return nil, fmt.Errorf("failed to save submission: %w", err)
	}

	log.Printf("💾 Saved submission ID: %d", submission.ID)

	// Step 4: Publish event to RabbitMQ (async, don't block response)
	go func() {
		event := model.SubmissionEvent{
			Event:         "SubmissionCompleted",
			UserID:        req.UserID,
			SkillTag:      skillTag,
			ScoreObtained: score,
			Timestamp:     time.Now().Format(time.RFC3339),
		}

		err := s.publisher.PublishSubmissionEvent(event)
		if err != nil {
			log.Printf("❌ Failed to publish event: %v", err)
		}
	}()

	return &model.SubmitResponse{
		Correct:  isCorrect,
		Score:    score,
		Feedback: feedback,
	}, nil
}
