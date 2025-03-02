package user

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

func (r *Repo) Upsert(ctx context.Context, input entity.User) (*entity.User, error) {
	// todo implement me
	panic("")
}

func (r *Repo) Read(ctx context.Context, filter entity.UserFilter) ([]entity.User, error) {
	// todo implement me
	panic("")
}

func (r *Repo) Delete(ctx context.Context, filter entity.UserFilter) error {
	// todo implement me
	panic("")
}

// todo should I use repo or send requests to auth
