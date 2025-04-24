package user

import (
	"context"

	"github.com/rashevskiivv/api/internal/entity"
)

type Repository interface {
	Insert(ctx context.Context, input entity.User) (*entity.User, error)
	Update(ctx context.Context, input entity.UserAuthInput) (*entity.User, error)
	Read(ctx context.Context, filter entity.UserFilter) ([]entity.User, error)
	Delete(ctx context.Context, filter entity.UserFilter) error
}
