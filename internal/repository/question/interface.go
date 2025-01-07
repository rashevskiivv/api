package question

import (
	"context"
	"tax-api/internal/entity"
)

type Repository interface {
	UpsertQuestion(ctx context.Context, input entity.Question) (*entity.Question, error)
	ReadQuestions(ctx context.Context, filter entity.QuestionFilter) ([]entity.Question, error)
	DeleteQuestion(ctx context.Context, filter entity.QuestionFilter) error
}
