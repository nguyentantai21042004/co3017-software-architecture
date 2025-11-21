package curl

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// LearnerServiceClient handles communication with Learner Service
type LearnerServiceClient struct {
	client *Client
}

// Ensure LearnerServiceClient implements LearnerServiceClientInterface
var _ LearnerServiceClientInterface = (*LearnerServiceClient)(nil)

// NewLearnerServiceClient creates a new Learner Service client
func NewLearnerServiceClient(baseURL string) *LearnerServiceClient {
	return &LearnerServiceClient{
		client: NewClient(baseURL, 10*time.Second),
	}
}

// GetMastery fetches mastery data for a user and skill
func (c *LearnerServiceClient) GetMastery(ctx context.Context, userID, skillTag string) (*MasteryResponse, error) {
	url := fmt.Sprintf("%s/internal/learner/%s/mastery?skill=%s", c.client.GetBaseURL(), userID, skillTag)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("%s: %w | %s=%s | %s=%s | error=%v",
			ErrMsgRequestCreationFailed, ErrServiceUnavailable, ErrCtxServiceName, "learner-service",
			ErrCtxURL, url, err)
	}

	resp, err := c.client.GetHTTPClient().Do(req)
	if err != nil {
		return nil, fmt.Errorf("%s: %w | %s=%s | %s=%s | %s=%s | error=%v",
			ErrMsgRequestExecutionFailed, ErrServiceUnavailable, ErrCtxServiceName, "learner-service",
			ErrCtxURL, url, ErrCtxUserID, userID, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("%s: %w | %s=%s | %s=%s | %s=%s | %s=%d",
			ErrMsgResourceNotFound, ErrNotFound, ErrCtxServiceName, "learner-service",
			ErrCtxUserID, userID, ErrCtxSkillTag, skillTag, ErrCtxStatusCode, resp.StatusCode)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s: %w | %s=%s | %s=%s | %s=%d",
			ErrMsgUnexpectedStatusCode, ErrServiceUnavailable, ErrCtxServiceName, "learner-service",
			ErrCtxURL, url, ErrCtxStatusCode, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%s: %w | %s=%s | %s=%s | error=%v",
			ErrMsgResponseReadFailed, ErrInvalidResponse, ErrCtxServiceName, "learner-service",
			ErrCtxURL, url, err)
	}

	var masteryResp MasteryResponse
	if err := json.Unmarshal(body, &masteryResp); err != nil {
		return nil, fmt.Errorf("%s: %w | %s=%s | %s=%s | %s=%s | error=%v",
			ErrMsgResponseParseFailed, ErrInvalidResponse, ErrCtxServiceName, "learner-service",
			ErrCtxURL, url, ErrCtxResponseBody, string(body), err)
	}

	return &masteryResp, nil
}
