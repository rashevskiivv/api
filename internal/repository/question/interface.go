package question

import (
	"context"
	"tax-api/internal/entity"
)

type Repository interface {
	Upsert(ctx context.Context, input entity.Question) (*entity.Question, error)
	Read(ctx context.Context, filter entity.QuestionFilter) ([]entity.Question, error)
	Delete(ctx context.Context, filter entity.QuestionFilter) error
}
