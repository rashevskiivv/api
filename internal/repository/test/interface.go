package test

import (
	"context"
	"tax-api/internal/entity"
)

type Repository interface {
	UpsertTest(ctx context.Context, input entity.Test) (*entity.Test, error)
	ReadTests(ctx context.Context, filter entity.TestFilter) ([]entity.Test, error)
	DeleteTest(ctx context.Context, filter entity.TestFilter) error
}
