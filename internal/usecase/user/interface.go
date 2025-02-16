package user

import (
	"context"
	"tax-api/internal/entity"
)

type UseCaseI interface {
	UpsertUser(ctx context.Context, input entity.User) (*entity.User, error)
	ReadUsers(ctx context.Context, input entity.UserFilter) ([]entity.User, error)
	DeleteUser(ctx context.Context, input entity.UserFilter) error
}
