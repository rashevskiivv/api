package vacancy

import (
	"context"
	"tax-api/internal/entity"
)

type Repository interface {
	Upsert(ctx context.Context, input entity.Vacancy) (*entity.Vacancy, error)
	Read(ctx context.Context, filter entity.VacancyFilter) ([]entity.Vacancy, error)
	Delete(ctx context.Context, filter entity.VacancyFilter) error
}
