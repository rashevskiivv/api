package test

import (
	"context"
	"tax-api/internal/entity"
)

type Repository interface {
	Upsert(ctx context.Context, input entity.Test) (*entity.Test, error)
	Read(ctx context.Context, filter entity.TestFilter) ([]entity.Test, error)
	Delete(ctx context.Context, filter entity.TestFilter) error

	Start(ctx context.Context, input entity.StartTestInput) (*entity.StartTestOutput, error)
	End(ctx context.Context, filter entity.EndTestInput) (*entity.EndTestOutput, error)
}
