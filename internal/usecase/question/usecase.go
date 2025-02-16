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
	log.Println("question usecase upsert started")

	output, err := uc.repo.Upsert(ctx, input)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("question usecase upsert done")

	return output, nil
}

func (uc *UseCase) ReadQuestions(ctx context.Context, input entity.QuestionFilter) ([]entity.Question, error) {
	log.Println("question usecase read started")
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
	log.Println("question usecase read done")

	return output, nil
}

func (uc *UseCase) DeleteQuestion(ctx context.Context, input entity.QuestionFilter) error {
	log.Println("question usecase delete started")
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
	log.Println("question usecase delete done")

	return nil
}
