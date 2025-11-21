package curl

import (
	"net/http"
	"time"
)

// Client is a wrapper around http.Client with common configurations
type Client struct {
	httpClient *http.Client
	baseURL    string
}

// NewClient creates a new HTTP client with timeout
func NewClient(baseURL string, timeout time.Duration) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: timeout,
		},
		baseURL: baseURL,
	}
}

// GetHTTPClient returns the underlying http.Client
func (c *Client) GetHTTPClient() *http.Client {
	return c.httpClient
}

// GetBaseURL returns the base URL
func (c *Client) GetBaseURL() string {
	return c.baseURL
}
