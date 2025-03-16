package link

import (
	"context"
	"tax-api/internal/entity"
)

type UseCaseI interface {
	UpsertAQ(ctx context.Context, input entity.AnswerQuestion) (*entity.AnswerQuestion, error)
	ReadAQ(ctx context.Context, filter entity.AnswerQuestionFilter) ([]entity.AnswerQuestion, error)
	DeleteAQ(ctx context.Context, filter entity.AnswerQuestionFilter) error

	UpsertQT(ctx context.Context, input entity.QuestionTest) (*entity.QuestionTest, error)
	ReadQT(ctx context.Context, filter entity.QuestionTestFilter) ([]entity.QuestionTest, error)
	DeleteQT(ctx context.Context, filter entity.QuestionTestFilter) error

	UpsertTS(ctx context.Context, input entity.TestSkill) (*entity.TestSkill, error)
	ReadTS(ctx context.Context, filter entity.TestSkillFilter) ([]entity.TestSkill, error)
	DeleteTS(ctx context.Context, filter entity.TestSkillFilter) error

	UpsertUS(ctx context.Context, input entity.UserSkill) (*entity.UserSkill, error)
	ReadUS(ctx context.Context, filter entity.UserSkillFilter) ([]entity.UserSkill, error)
	DeleteUS(ctx context.Context, filter entity.UserSkillFilter) error

	UpsertSV(ctx context.Context, input entity.SkillVacancy) (*entity.SkillVacancy, error)
	ReadSV(ctx context.Context, filter entity.SkillVacancyFilter) ([]entity.SkillVacancy, error)
	DeleteSV(ctx context.Context, filter entity.SkillVacancyFilter) error
}
