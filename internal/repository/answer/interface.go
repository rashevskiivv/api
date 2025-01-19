package answer

import (
	"context"
	"tax-api/internal/entity"
)

type Repository interface {
	UpsertAnswer(ctx context.Context, input entity.Answer) (*entity.Answer, error)
	ReadAnswers(ctx context.Context, filter entity.AnswerFilter) ([]entity.Answer, error)
	DeleteAnswer(ctx context.Context, filter entity.AnswerFilter) error
}
