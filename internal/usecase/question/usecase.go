package question

import (
	"context"
	"log"
	"tax-api/internal/entity"
	repositoryQuestion "tax-api/internal/repository/question"
)

type UseCase struct {
	repo repositoryQuestion.Repository
}

func NewUseCase(repo repositoryQuestion.Repository) *UseCase {
	return &UseCase{repo: repo}
}

func (uc *UseCase) UpsertQuestion(ctx context.Context, input entity.Question) (*entity.Question, error) {
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

func (uc *UseCase) ReadQuestions(ctx context.Context, input entity.QuestionFilter) ([]entity.Question, error) {
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

func (uc *UseCase) DeleteQuestion(ctx context.Context, input entity.QuestionFilter) error {
	err := input.Validate()
	if err != nil {
		log.Println(err)
		return err
	}

	return uc.repo.Delete(ctx, input)
}
