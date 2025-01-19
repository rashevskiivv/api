package vacancy

import (
	"context"
	"tax-api/internal/entity"
)

type Repository interface {
	UpsertVacancy(ctx context.Context, input entity.Vacancy) (*entity.Vacancy, error)
	ReadVacancies(ctx context.Context, filter entity.VacancyFilter) ([]entity.Vacancy, error)
	DeleteVacancy(ctx context.Context, filter entity.VacancyFilter) error
}
