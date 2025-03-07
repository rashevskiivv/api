package skill

import (
	"context"
	"tax-api/internal/entity"
)

type Repository interface {
	Upsert(ctx context.Context, input entity.Skill) (*entity.Skill, error)
	Read(ctx context.Context, filter entity.SkillFilter) ([]entity.Skill, error)
	Delete(ctx context.Context, filter entity.SkillFilter) error
}
