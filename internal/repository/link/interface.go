package link

import (
	"context"
	"tax-api/internal/entity"
)

type Repository interface {
	UpsertAnswerQuestion(ctx context.Context, input entity.AnswerQuestion) (*entity.AnswerQuestion, error)
	ReadAnswerQuestion(ctx context.Context, filter entity.AnswerQuestionFilter) ([]entity.AnswerQuestion, error)
	DeleteAnswerQuestion(ctx context.Context, filter entity.AnswerQuestionFilter) error

	UpsertQuestionTest(ctx context.Context, input entity.QuestionTest) (*entity.QuestionTest, error)
	ReadQuestionTest(ctx context.Context, filter entity.QuestionTestFilter) ([]entity.QuestionTest, error)
	DeleteQuestionTest(ctx context.Context, filter entity.QuestionTestFilter) error

	UpsertTestSkill(ctx context.Context, input entity.TestSkill) (*entity.TestSkill, error)
	ReadTestSkill(ctx context.Context, filter entity.TestSkillFilter) ([]entity.TestSkill, error)
	DeleteTestSkill(ctx context.Context, filter entity.TestSkillFilter) error

	UpsertUserSkill(ctx context.Context, input entity.UserSkill) (*entity.UserSkill, error)
	ReadUserSkill(ctx context.Context, filter entity.UserSkillFilter) ([]entity.UserSkill, error)
	DeleteUserSkill(ctx context.Context, filter entity.UserSkillFilter) error

	UpsertSkillVacancy(ctx context.Context, input entity.SkillVacancy) (*entity.SkillVacancy, error)
	ReadSkillVacancy(ctx context.Context, filter entity.SkillVacancyFilter) ([]entity.SkillVacancy, error)
	DeleteSkillVacancy(ctx context.Context, filter entity.SkillVacancyFilter) error
}
