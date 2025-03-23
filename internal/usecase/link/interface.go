package link

import (
	"context"
	"tax-api/internal/entity"
)

type UseCaseI interface {
	UpsertTS(ctx context.Context, input entity.TestSkill) error
	ReadTS(ctx context.Context, filter entity.TestSkillFilter) ([]entity.TestSkill, error)
	DeleteTS(ctx context.Context, filter entity.TestSkillFilter) error

	UpsertUS(ctx context.Context, input entity.UserSkill) error
	ReadUS(ctx context.Context, filter entity.UserSkillFilter) ([]entity.UserSkill, error)
	DeleteUS(ctx context.Context, filter entity.UserSkillFilter) error

	UpsertSV(ctx context.Context, input entity.SkillVacancy) error
	ReadSV(ctx context.Context, filter entity.SkillVacancyFilter) ([]entity.SkillVacancy, error)
	DeleteSV(ctx context.Context, filter entity.SkillVacancyFilter) error
}
