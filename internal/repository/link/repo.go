package link

import (
	"context"
	"tax-api/internal/entity"
	"tax-api/internal/repository"

	"github.com/Masterminds/squirrel"
)

type Repo struct {
	*repository.Postgres
	builder squirrel.StatementBuilderType
}

func NewRepo(pg *repository.Postgres) *Repo {
	return &Repo{
		Postgres: pg,
		builder:  squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (r *Repo) UpsertAnswerQuestion(ctx context.Context, input entity.AnswerQuestion) (*entity.AnswerQuestion, error) {

}

func (r *Repo) ReadAnswerQuestion(ctx context.Context, filter entity.AnswerQuestionFilter) ([]entity.AnswerQuestion, error) {

}

func (r *Repo) DeleteAnswerQuestion(ctx context.Context, filter entity.AnswerQuestionFilter) error {

}

func (r *Repo) UpsertQuestionTest(ctx context.Context, input entity.QuestionTest) (*entity.QuestionTest, error) {

}

func (r *Repo) ReadQuestionTest(ctx context.Context, filter entity.QuestionTestFilter) ([]entity.QuestionTest, error) {

}

func (r *Repo) DeleteQuestionTest(ctx context.Context, filter entity.QuestionTestFilter) error {

}

func (r *Repo) UpsertTestSkill(ctx context.Context, input entity.TestSkill) (*entity.TestSkill, error) {

}

func (r *Repo) ReadTestSkill(ctx context.Context, filter entity.TestSkillFilter) ([]entity.TestSkill, error) {

}

func (r *Repo) DeleteTestSkill(ctx context.Context, filter entity.TestSkillFilter) error {

}

func (r *Repo) UpsertUserSkill(ctx context.Context, input entity.UserSkill) (*entity.UserSkill, error) {

}

func (r *Repo) ReadUserSkill(ctx context.Context, filter entity.UserSkillFilter) ([]entity.UserSkill, error) {

}

func (r *Repo) DeleteUserSkill(ctx context.Context, filter entity.UserSkillFilter) error {

}

func (r *Repo) UpsertSkillVacancy(ctx context.Context, input entity.SkillVacancy) (*entity.SkillVacancy, error) {

}

func (r *Repo) ReadSkillVacancy(ctx context.Context, filter entity.SkillVacancyFilter) ([]entity.SkillVacancy, error) {

}

func (r *Repo) DeleteSkillVacancy(ctx context.Context, filter entity.SkillVacancyFilter) error {

}
