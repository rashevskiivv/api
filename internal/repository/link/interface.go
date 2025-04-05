package link

import (
	"context"

	"github.com/rashevskiivv/api/internal/entity"
)

type Repository interface {
	UpsertTestSkill(ctx context.Context, input entity.TestSkill) error
	DeleteTestSkill(ctx context.Context, input entity.TestSkillFilter) error

	UpsertUserSkill(ctx context.Context, input entity.UserSkill) error
	ReadUserSkill(ctx context.Context, input entity.UserSkillFilter) ([]entity.UserSkill, error)
	DeleteUserSkill(ctx context.Context, input entity.UserSkillFilter) error

	UpsertSkillVacancy(ctx context.Context, input entity.SkillVacancy) error
	ReadSkillVacancy(ctx context.Context, input entity.SkillVacancyFilter) ([]entity.SkillVacancy, error)
	DeleteSkillVacancy(ctx context.Context, input entity.SkillVacancyFilter) error
}
