package test

import (
	"context"
	"fmt"
	"log"

	"github.com/rashevskiivv/api/internal/entity"
	repositoryAnswer "github.com/rashevskiivv/api/internal/repository/answer"
	repositoryQuestion "github.com/rashevskiivv/api/internal/repository/question"
	repositoryTest "github.com/rashevskiivv/api/internal/repository/test"
)

type UseCase struct {
	repo          repositoryTest.Repository
	repoQuestions repositoryQuestion.Repository
	repoAnswers   repositoryAnswer.Repository
}

func NewUseCase(repo repositoryTest.Repository, repoQuestions repositoryQuestion.Repository, repoAnswers repositoryAnswer.Repository) *UseCase {
	return &UseCase{
		repo:          repo,
		repoQuestions: repoQuestions,
		repoAnswers:   repoAnswers,
	}
}

func (uc *UseCase) UpsertTest(ctx context.Context, input entity.Test) (*entity.Test, error) {
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

func (uc *UseCase) ReadTests(ctx context.Context, input entity.TestFilter) ([]entity.Test, error) {
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

func (uc *UseCase) DeleteTest(ctx context.Context, input entity.TestFilter) error {
	err := input.Validate()
	if err != nil {
		log.Println(err)
		return err
	}

	return uc.repo.Delete(ctx, input)
}

func (uc *UseCase) StartTest(ctx context.Context, input entity.StartTestInput) (*entity.StartTestOutput, error) {
	/* todo what for i require id_user? make another table with stats? */
	err := input.Validate()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	testInput := entity.TestFilter{
		ID:    []int64{input.IDTest},
		Limit: 1,
	}
	testOutput, err := uc.repo.Read(ctx, testInput)
	if err != nil {
		return nil, err
	}
	if len(testOutput) == 0 {
		return nil, fmt.Errorf("no tests found")
	}

	questionsInput := entity.QuestionFilter{
		IDTest: []int64{*testOutput[0].ID},
	}
	questionsOutput, err := uc.repoQuestions.Read(ctx, questionsInput)
	if err != nil {
		return nil, err
	}
	if len(questionsOutput) == 0 {
		return nil, fmt.Errorf("no questions found")
	}

	var questionsToReturn []entity.QuestionToReturn
	for _, question := range questionsOutput {
		answersToReturn := make([]entity.AnswerToReturn, 0, 4)
		answersInput := entity.AnswerFilter{
			IDQuestion: []int64{*question.ID},
		}
		answersOutput, err := uc.repoAnswers.Read(ctx, answersInput)
		if err != nil {
			return nil, err
		}
		if len(answersOutput) == 0 {
			return nil, fmt.Errorf("no answers found for question #%v", question.ID)
		}

		for _, answer := range answersOutput {
			answerToReturn := entity.AnswerToReturn{
				ID:      *answer.ID,
				Answer:  answer.Answer,
				IsRight: answer.IsRight,
			}
			answersToReturn = append(answersToReturn, answerToReturn)
		}

		questionToReturn := entity.QuestionToReturn{
			ID:       *question.ID,
			Question: question.Question,
			Answers:  answersToReturn,
		}
		questionsToReturn = append(questionsToReturn, questionToReturn)
	}

	testToReturn := entity.StartTestOutput{
		NumberOfQuestions: int8(len(questionsToReturn)),
		Questions:         questionsToReturn,
	}
	return &testToReturn, nil
}

func (uc *UseCase) EndTest(ctx context.Context, filter entity.EndTestInput) error {
	// todo implement me
	panic("implement me")
}
