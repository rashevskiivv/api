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

	output, err := uc.repo.Upsert(ctx, input)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("vacancy usecase upsert done")

	return output, nil
}

func (uc *UseCase) ReadVacancies(ctx context.Context, input entity.VacancyFilter) ([]entity.Vacancy, error) {
	log.Println("vacancy usecase read started")
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
	log.Println("vacancy usecase read done")

	return output, nil
}

func (uc *UseCase) DeleteVacancy(ctx context.Context, input entity.VacancyFilter) error {
	log.Println("vacancy usecase delete started")
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
	log.Println("vacancy usecase delete done")

	return nil
}
