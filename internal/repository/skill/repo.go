package skill

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

func (r *Repo) Upsert(ctx context.Context, input entity.Skill) (*entity.Skill, error) {
	// todo
	panic("implement me")
}

func (r *Repo) Read(ctx context.Context, filter entity.SkillFilter) ([]entity.Skill, error) {
	// todo
	panic("implement me")
}

func (r *Repo) Delete(ctx context.Context, filter entity.SkillFilter) error {
	// todo
	panic("implement me")
}
