package curl

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// ContentServiceClient handles communication with Content Service
type ContentServiceClient struct {
	client *Client
}

// Ensure ContentServiceClient implements ContentServiceClientInterface
var _ ContentServiceClientInterface = (*ContentServiceClient)(nil)

// NewContentServiceClient creates a new Content Service client
func NewContentServiceClient(baseURL string) *ContentServiceClient {
	return &ContentServiceClient{
		client: NewClient(baseURL, 10*time.Second),
	}
}

// GetRecommendation fetches content recommendation based on skill and type
func (c *ContentServiceClient) GetRecommendation(ctx context.Context, skillTag, contentType, userID string) (*ContentResponse, error) {
	url := fmt.Sprintf("%s/api/content/recommend?skill=%s&type=%s&userId=%s", c.client.GetBaseURL(), skillTag, contentType, userID)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("%s: %w | %s=%s | %s=%s | %s=%s | %s=%s | error=%v",
			ErrMsgRequestCreationFailed, ErrServiceUnavailable, ErrCtxServiceName, "content-service",
			ErrCtxURL, url, ErrCtxSkillTag, skillTag, ErrCtxContentType, contentType, err)
	}

	resp, err := c.client.GetHTTPClient().Do(req)
	if err != nil {
		return nil, fmt.Errorf("%s: %w | %s=%s | %s=%s | %s=%s | %s=%s | error=%v",
			ErrMsgRequestExecutionFailed, ErrServiceUnavailable, ErrCtxServiceName, "content-service",
			ErrCtxURL, url, ErrCtxSkillTag, skillTag, ErrCtxContentType, contentType, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("%s: %w | %s=%s | %s=%s | %s=%s | %s=%d",
			ErrMsgResourceNotFound, ErrNotFound, ErrCtxServiceName, "content-service",
			ErrCtxSkillTag, skillTag, ErrCtxContentType, contentType, ErrCtxStatusCode, resp.StatusCode)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s: %w | %s=%s | %s=%s | %s=%d",
			ErrMsgUnexpectedStatusCode, ErrServiceUnavailable, ErrCtxServiceName, "content-service",
			ErrCtxURL, url, ErrCtxStatusCode, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%s: %w | %s=%s | %s=%s | error=%v",
			ErrMsgResponseReadFailed, ErrInvalidResponse, ErrCtxServiceName, "content-service",
			ErrCtxURL, url, err)
	}

	var contentResp ContentResponse
	if err := json.Unmarshal(body, &contentResp); err != nil {
		return nil, fmt.Errorf("%s: %w | %s=%s | %s=%s | %s=%s | error=%v",
			ErrMsgResponseParseFailed, ErrInvalidResponse, ErrCtxServiceName, "content-service",
			ErrCtxURL, url, ErrCtxResponseBody, string(body), err)
	}

	return &contentResp, nil
}
