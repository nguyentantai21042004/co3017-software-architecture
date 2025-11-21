package adaptive

import "context"

//go:generate mockery --name UseCase
type UseCase interface {
	RecommendNextLesson(ctx context.Context, input RecommendInput) (RecommendOutput, error)
}
