package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"adaptive-engine-service/internal/model"
)

const (
	MASTERY_THRESHOLD = 50
)

type AdaptiveService interface {
	RecommendNextLesson(req *model.NextLessonRequest) (*model.NextLessonResponse, error)
}

type adaptiveService struct {
	learnerServiceURL string
	contentServiceURL string
}

func NewAdaptiveService(learnerURL, contentURL string) AdaptiveService {
	return &adaptiveService{
		learnerServiceURL: learnerURL,
		contentServiceURL: contentURL,
	}
}

func (s *adaptiveService) RecommendNextLesson(req *model.NextLessonRequest) (*model.NextLessonResponse, error) {
	log.Printf("🧠 Adaptive Engine: user=%s, skill=%s", req.UserID, req.CurrentSkill)

	mastery, err := s.fetchMastery(req.UserID, req.CurrentSkill)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch mastery: %w", err)
	}

	var contentType, reason string
	if mastery.MasteryScore < MASTERY_THRESHOLD {
		contentType = "remedial"
		reason = fmt.Sprintf("Your mastery is %d%%. Let's review the basics.", mastery.MasteryScore)
		log.Printf("🔄 Recommending REMEDIAL (score=%d < %d)", mastery.MasteryScore, MASTERY_THRESHOLD)
	} else {
		contentType = "standard"
		reason = fmt.Sprintf("Great! Your mastery is %d%%. Continue with the next challenge.", mastery.MasteryScore)
		log.Printf("✅ Recommending STANDARD (score=%d >= %d)", mastery.MasteryScore, MASTERY_THRESHOLD)
	}

	question, err := s.fetchContent(req.CurrentSkill, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch content: %w", err)
	}

	log.Printf("📚 Recommended question ID: %d (type: %s)", question.Data.ID, contentType)

	return &model.NextLessonResponse{
		NextLessonID: int(question.Data.ID),
		Reason:       reason,
		MasteryScore: mastery.MasteryScore,
		ContentType:  contentType,
	}, nil
}

func (s *adaptiveService) fetchMastery(userID, skillTag string) (*model.MasteryResponse, error) {
	url := fmt.Sprintf("%s/internal/learner/%s/mastery?skill=%s", s.learnerServiceURL, userID, skillTag)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var mastery model.MasteryResponse
	json.Unmarshal(body, &mastery)
	return &mastery, nil
}

func (s *adaptiveService) fetchContent(skillTag, contentType string) (*model.QuestionResponse, error) {
	url := fmt.Sprintf("%s/api/content/recommend?skill=%s&type=%s", s.contentServiceURL, skillTag, contentType)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var question model.QuestionResponse
	json.Unmarshal(body, &question)
	return &question, nil
}
