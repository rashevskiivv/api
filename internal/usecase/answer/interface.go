package answer

import (
	"context"

	"github.com/rashevskiivv/api/internal/entity"
)

type UseCaseI interface {
	UpsertAnswer(ctx context.Context, input entity.Answer) (*entity.Answer, error)
	ReadAnswers(ctx context.Context, input entity.AnswerFilter) ([]entity.Answer, error)
	DeleteAnswer(ctx context.Context, input entity.AnswerFilter) error
}
