package vacancy

import (
	"context"

	"github.com/rashevskiivv/api/internal/entity"
)

type UseCaseI interface {
	CloseIdleConnections()
	UpsertVacancy(ctx context.Context, input entity.VacancyInput) (*entity.Vacancy, error)
	ReadVacancies(ctx context.Context, input entity.VacancyFilter) ([]entity.Vacancy, error)
	DeleteVacancy(ctx context.Context, input entity.VacancyFilter) error
}
