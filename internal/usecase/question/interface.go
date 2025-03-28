package question

import (
	"context"
	"tax-api/internal/entity"
)

type UseCaseI interface {
	UpsertQuestion(ctx context.Context, input entity.Question) (*entity.Question, error)
	ReadQuestions(ctx context.Context, input entity.QuestionFilter) ([]entity.Question, error)
	DeleteQuestion(ctx context.Context, input entity.QuestionFilter) error
}
