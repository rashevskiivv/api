package answer

import (
	"context"
	"log"
	"tax-api/internal/entity"
	repositoryAnswer "tax-api/internal/repository/answer"
)

type UseCase struct {
	repo repositoryAnswer.Repository
}

func NewUseCase(repo repositoryAnswer.Repository) *UseCase {
	return &UseCase{repo: repo}
}

func (uc *UseCase) UpsertAnswer(ctx context.Context, input entity.Answer) (*entity.Answer, error) {
	log.Println("answer usecase upsert started")

	output, err := uc.repo.Upsert(ctx, input)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("answer usecase upsert done")

	return output, nil
}

func (uc *UseCase) ReadAnswers(ctx context.Context, input entity.AnswerFilter) ([]entity.Answer, error) {
	log.Println("answer usecase read started")
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
	log.Println("answer usecase read done")

	return output, nil
}

func (uc *UseCase) DeleteAnswer(ctx context.Context, input entity.AnswerFilter) error {
	log.Println("answer usecase delete started")
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
	log.Println("answer usecase delete done")

	return nil
}
