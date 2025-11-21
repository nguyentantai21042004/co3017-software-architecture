package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// ContentClient defines the interface for interacting with the Content Service
type ContentClient interface {
	FetchQuestion(questionID int64) (string, string, error)
}

// HttpContentClient implements ContentClient using HTTP
type HttpContentClient struct {
	BaseURL string
}

func NewHttpContentClient(baseURL string) *HttpContentClient {
	return &HttpContentClient{BaseURL: baseURL}
}

func (c *HttpContentClient) FetchQuestion(questionID int64) (string, string, error) {
	url := fmt.Sprintf("%s/%d", c.BaseURL, questionID)

	log.Printf("🔍 Fetching question from Content Service: %s", url)

	resp, err := http.Get(url)
	if err != nil {
		return "", "", fmt.Errorf("failed to call content service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("content service returned status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", fmt.Errorf("failed to read response body: %w", err)
	}

	var contentResp ContentServiceResponse
	err = json.Unmarshal(body, &contentResp)
	if err != nil {
		return "", "", fmt.Errorf("failed to parse response: %w", err)
	}

	if !contentResp.Success {
		return "", "", fmt.Errorf("content service error: %s", contentResp.Message)
	}

	log.Printf("✅ Got correct answer: %s, skill: %s",
		contentResp.Data.CorrectAnswer, contentResp.Data.SkillTag)

	return contentResp.Data.CorrectAnswer, contentResp.Data.SkillTag, nil
}
