package curl

import "context"

// ContentServiceClientInterface defines the interface for content service client
type ContentServiceClientInterface interface {
	GetQuestion(ctx context.Context, questionID int64) (*ContentQuestionResponse, error)
}

// Ensure ContentServiceClient implements ContentServiceClientInterface
var _ ContentServiceClientInterface = (*ContentServiceClient)(nil)
