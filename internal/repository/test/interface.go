package test

import (
	"context"

	"github.com/rashevskiivv/api/internal/entity"
)

type Repository interface {
	Upsert(ctx context.Context, input entity.Test) (*entity.Test, error)
	Read(ctx context.Context, filter entity.TestFilter) ([]entity.Test, error)
	Delete(ctx context.Context, filter entity.TestFilter) error
}
