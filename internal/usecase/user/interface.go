package user

import (
	"context"

	"github.com/rashevskiivv/api/internal/entity"
)

type UseCaseI interface {
	CloseIdleConnections()
	UpsertUser(input entity.UserAuthInput) (*entity.User, error)
	ReadUsers(ctx context.Context, input entity.UserFilter) ([]entity.User, error)
	DeleteUser(ctx context.Context, input entity.UserFilter) error
}
