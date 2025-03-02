package test

import (
	"context"
	"tax-api/internal/entity"
)

type UseCaseI interface {
	UpsertTest(ctx context.Context, input entity.Test) (*entity.Test, error)
	ReadTests(ctx context.Context, input entity.TestFilter) ([]entity.Test, error)
	DeleteTest(ctx context.Context, input entity.TestFilter) error
}
