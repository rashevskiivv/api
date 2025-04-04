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

func (uc *UseCase) ReadAnswers(ctx context.Context, input entity.AnswerFilter) ([]entity.Answer, error) {
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

func (uc *UseCase) DeleteAnswer(ctx context.Context, input entity.AnswerFilter) error {
	err := input.Validate()
	if err != nil {
		log.Println(err)
		return err
	}

	return uc.repo.Delete(ctx, input)
}
