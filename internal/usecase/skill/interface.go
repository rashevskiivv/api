package skill

import (
	"context"
	"tax-api/internal/entity"
)

type UseCaseI interface {
	UpsertSkill(ctx context.Context, input entity.Skill) (*entity.Skill, error)
	ReadSkills(ctx context.Context, input entity.SkillFilter) ([]entity.Skill, error)
	DeleteSkill(ctx context.Context, input entity.SkillFilter) error
}
