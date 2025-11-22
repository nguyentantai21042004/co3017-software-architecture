package curl

import "context"

// LearnerServiceClientInterface defines the interface for learner service client
type LearnerServiceClientInterface interface {
	GetMastery(ctx context.Context, userID, skillTag string) (*MasteryResponse, error)
}

// ContentServiceClientInterface defines the interface for content service client
type ContentServiceClientInterface interface {
	GetRecommendation(ctx context.Context, skillTag, contentType, userID string) (*ContentResponse, error)
}
