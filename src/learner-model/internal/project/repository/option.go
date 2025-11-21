package repository

import (
	"learner-model-service/internal/model"
	"learner-model-service/pkg/paginator"
)

type CreateOptions struct {
	Project model.Project
}

type UpdateOptions struct {
	Project model.Project
}

type GetOptions struct {
	IDs           []string
	Statuses      []string
	CreatedBy     *string
	SearchName    *string
	PaginateQuery paginator.PaginateQuery
}

type ListOptions struct {
	IDs        []string
	Statuses   []string
	CreatedBy  *string
	SearchName *string
}

type GetOneOptions struct {
	ID string
}
