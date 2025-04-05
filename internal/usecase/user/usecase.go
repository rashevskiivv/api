package user

import (
	"context"
	"log"

	"github.com/rashevskiivv/api/internal/entity"
	repositoryUser "github.com/rashevskiivv/api/internal/repository/user"
)

type UseCase struct {
	repo repositoryUser.Repository
}

func NewUseCase(repo repositoryUser.Repository) *UseCase {
	return &UseCase{repo: repo}
}

// todo should I use repo or send requests to auth

func (uc *UseCase) UpsertUser(ctx context.Context, input entity.User) (*entity.User, error) {
	output, err := uc.repo.Upsert(ctx, input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (uc *UseCase) ReadUsers(ctx context.Context, input entity.UserFilter) ([]entity.User, error) {
	err := input.Validate()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	output, err := uc.repo.Read(ctx, input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (uc *UseCase) DeleteUser(ctx context.Context, input entity.UserFilter) error {
	err := input.Validate()
	if err != nil {
		log.Println(err)
		return err
	}

	return uc.repo.Delete(ctx, input)
}
