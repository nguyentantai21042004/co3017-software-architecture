package scoring

import "context"

// UseCase defines the interface for scoring use cases
//
//go:generate mockery --name UseCase
type UseCase interface {
	SubmitAnswer(ctx context.Context, input SubmitInput) (SubmitOutput, error)
	GetAnsweredQuestions(ctx context.Context, userID, skillTag string) ([]int64, error)
}
