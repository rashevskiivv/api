package test

import (
	"context"
	"log"
	"tax-api/internal/entity"
	repositoryTest "tax-api/internal/repository/test"
)

type UseCase struct {
	repo repositoryTest.Repository
}

func NewUseCase(repo repositoryTest.Repository) *UseCase {
	return &UseCase{repo: repo}
}

func (uc *UseCase) UpsertTest(ctx context.Context, input entity.Test) (*entity.Test, error) {
	log.Println("test usecase upsert started")

	output, err := uc.repo.Upsert(ctx, input)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("test usecase upsert done")

	return output, nil
}

func (uc *UseCase) ReadTests(ctx context.Context, input entity.TestFilter) ([]entity.Test, error) {
	log.Println("test usecase read started")
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
	log.Println("test usecase read done")

	return output, nil
}

func (uc *UseCase) DeleteTest(ctx context.Context, input entity.TestFilter) error {
	log.Println("test usecase delete started")
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
	log.Println("test usecase delete done")

	return nil
}
