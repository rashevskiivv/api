package answer

import (
	"context"

	"github.com/rashevskiivv/api/internal/entity"
)

type Repository interface {
	Upsert(ctx context.Context, input entity.Answer) (*entity.Answer, error)
	Read(ctx context.Context, filter entity.AnswerFilter) ([]entity.Answer, error)
	Delete(ctx context.Context, filter entity.AnswerFilter) error
}
