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
	log.Println("vacancy usecase upsert started")
	defer log.Println("vacancy usecase upsert done")

	err := input.Validate()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	output, err := uc.repo.Upsert(ctx, input)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return output, nil
}

func (uc *UseCase) ReadVacancies(ctx context.Context, input entity.VacancyFilter) ([]entity.Vacancy, error) {
	log.Println("vacancy usecase read started")
	defer log.Println("vacancy usecase read done")

	err := input.Validate()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	output, err := uc.repo.Read(ctx, input)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return output, nil
}

func (uc *UseCase) DeleteVacancy(ctx context.Context, input entity.VacancyFilter) error {
	log.Println("vacancy usecase delete started")
	defer log.Println("vacancy usecase delete done")

	err := input.Validate()
	if err != nil {
		log.Println(err)
		return err
	}

	err = uc.repo.Delete(ctx, input)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
