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

// NewContentServiceClient creates a new Content Service client
func NewContentServiceClient(baseURL string) *ContentServiceClient {
	return &ContentServiceClient{
		client: NewClient(baseURL, 10*time.Second),
	}
}

// GetQuestion fetches question details by ID
func (c *ContentServiceClient) GetQuestion(ctx context.Context, questionID int64) (*ContentQuestionResponse, error) {
	url := fmt.Sprintf("%s/%d", c.client.GetBaseURL(), questionID)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("%s: %w | %s=%s | %s=%s | %s=%d | error=%v",
			ErrMsgRequestCreationFailed, ErrServiceUnavailable, ErrCtxServiceName, "content-service",
			ErrCtxURL, url, ErrCtxQuestionID, questionID, err)
	}

	resp, err := c.client.GetHTTPClient().Do(req)
	if err != nil {
		return nil, fmt.Errorf("%s: %w | %s=%s | %s=%s | %s=%d | error=%v",
			ErrMsgRequestExecutionFailed, ErrServiceUnavailable, ErrCtxServiceName, "content-service",
			ErrCtxURL, url, ErrCtxQuestionID, questionID, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("%s: %w | %s=%s | %s=%d | %s=%d",
			ErrMsgResourceNotFound, ErrNotFound, ErrCtxServiceName, "content-service",
			ErrCtxQuestionID, questionID, ErrCtxStatusCode, resp.StatusCode)
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

	var contentResp ContentQuestionResponse
	if err := json.Unmarshal(body, &contentResp); err != nil {
		return nil, fmt.Errorf("%s: %w | %s=%s | %s=%s | %s=%s | error=%v",
			ErrMsgResponseParseFailed, ErrInvalidResponse, ErrCtxServiceName, "content-service",
			ErrCtxURL, url, ErrCtxResponseBody, string(body), err)
	}

	if contentResp.ErrorCode != 0 {
		return nil, fmt.Errorf("content service error: %s | %s=%s | %s=%d",
			contentResp.Message, ErrCtxServiceName, "content-service", ErrCtxQuestionID, questionID)
	}

	return &contentResp, nil
}
