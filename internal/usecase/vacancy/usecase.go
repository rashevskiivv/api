package vacancy

import (
	"context"
	"log"
	"tax-api/internal/entity"
	repositoryVacancy "tax-api/internal/repository/vacancy"
)

type UseCase struct {
	repo repositoryVacancy.Repository
}

func NewUseCase(repo repositoryVacancy.Repository) *UseCase {
	return &UseCase{repo: repo}
}

func (uc *UseCase) UpsertVacancy(ctx context.Context, input entity.Vacancy) (*entity.Vacancy, error) {
	err := input.Validate()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	output, err := uc.repo.Upsert(ctx, input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (uc *UseCase) ReadVacancies(ctx context.Context, input entity.VacancyFilter) ([]entity.Vacancy, error) {
	err := input.Validate()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	output, err := uc.repo.Read(ctx, input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (uc *UseCase) DeleteVacancy(ctx context.Context, input entity.VacancyFilter) error {
	err := input.Validate()
	if err != nil {
		log.Println(err)
		return err
	}

	return uc.repo.Delete(ctx, input)
}
