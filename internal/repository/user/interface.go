package user

import (
	"context"
	"tax-api/internal/entity"
)

type Repository interface {
	Upsert(ctx context.Context, input entity.User) (*entity.User, error)
	Read(ctx context.Context, filter entity.UserFilter) ([]entity.User, error)
	Delete(ctx context.Context, filter entity.UserFilter) error
}
