package learner

import "context"

//go:generate mockery --name UseCase
type UseCase interface {
	GetMastery(ctx context.Context, input GetMasteryInput) (MasteryOutput, error)
	UpdateMasteryFromEvent(ctx context.Context, input UpdateMasteryInput) error
}
