package vacancy

import (
	"context"

	"github.com/rashevskiivv/api/internal/entity"
)

type Repository interface {
	Upsert(ctx context.Context, input entity.Vacancy) (*entity.Vacancy, error)
	Read(ctx context.Context, filter entity.VacancyFilter) ([]entity.Vacancy, error)
	Delete(ctx context.Context, filter entity.VacancyFilter) error
}
