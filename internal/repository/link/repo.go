package link

import (
	"github.com/rashevskiivv/api/internal/repository"

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
