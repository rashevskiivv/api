package link

import (
	"context"
	"tax-api/internal/entity"
	repositoryLink "tax-api/internal/repository/link"
)

type UseCase struct {
	repo repositoryLink.Repository
}

func NewUseCase(repo repositoryLink.Repository) *UseCase {
	return &UseCase{repo: repo}
}

func (uc *UseCase) UpsertAQ(ctx context.Context, input entity.AnswerQuestion) (*entity.AnswerQuestion, error) {

}

func (uc *UseCase) ReadAQ(ctx context.Context, filter entity.AnswerQuestionFilter) ([]entity.AnswerQuestion, error) {

}

func (uc *UseCase) DeleteAQ(ctx context.Context, filter entity.AnswerQuestionFilter) error {

}

func (uc *UseCase) UpsertQT(ctx context.Context, input entity.QuestionTest) (*entity.QuestionTest, error) {

}

func (uc *UseCase) ReadQT(ctx context.Context, filter entity.QuestionTestFilter) ([]entity.QuestionTest, error) {

}

func (uc *UseCase) DeleteQT(ctx context.Context, filter entity.QuestionTestFilter) error {

}

func (uc *UseCase) UpsertTS(ctx context.Context, input entity.TestSkill) (*entity.TestSkill, error) {

}

func (uc *UseCase) ReadTS(ctx context.Context, filter entity.TestSkillFilter) ([]entity.TestSkill, error) {

}

func (uc *UseCase) DeleteTS(ctx context.Context, filter entity.TestSkillFilter) error {

}

func (uc *UseCase) UpsertUS(ctx context.Context, input entity.UserSkill) (*entity.UserSkill, error) {

}

func (uc *UseCase) ReadUS(ctx context.Context, filter entity.UserSkillFilter) ([]entity.UserSkill, error) {

}

func (uc *UseCase) DeleteUS(ctx context.Context, filter entity.UserSkillFilter) error {

}

func (uc *UseCase) UpsertSV(ctx context.Context, input entity.SkillVacancy) (*entity.SkillVacancy, error) {

}

func (uc *UseCase) ReadSV(ctx context.Context, filter entity.SkillVacancyFilter) ([]entity.SkillVacancy, error) {

}

func (uc *UseCase) DeleteSV(ctx context.Context, filter entity.SkillVacancyFilter) error {

}
