package vacancy

import (
	"context"
	"tax-api/internal/entity"
)

type UseCaseI interface {
	UpsertVacancy(ctx context.Context, input entity.Vacancy) (*entity.Vacancy, error)
	ReadVacancies(ctx context.Context, input entity.VacancyFilter) ([]entity.Vacancy, error)
	DeleteVacancy(ctx context.Context, input entity.VacancyFilter) error
}
