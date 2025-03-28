package user

import (
	"context"
	"log"
	"tax-api/internal/entity"
	repositoryUser "tax-api/internal/repository/user"
)

type UseCase struct {
	repo repositoryUser.Repository
}

func NewUseCase(repo repositoryUser.Repository) *UseCase {
	return &UseCase{repo: repo}
}

// todo should I use repo or send requests to auth

func (uc *UseCase) UpsertUser(ctx context.Context, input entity.User) (*entity.User, error) {
	log.Println("user usecase upsert started")
	defer log.Println("user usecase upsert done")

	output, err := uc.repo.Upsert(ctx, input)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return output, nil
}

func (uc *UseCase) ReadUsers(ctx context.Context, input entity.UserFilter) ([]entity.User, error) {
	log.Println("user usecase read started")
	defer log.Println("user usecase read done")
	err := input.Validate()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	output, err := uc.repo.Read(ctx, input)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return output, nil
}

func (uc *UseCase) DeleteUser(ctx context.Context, input entity.UserFilter) error {
	log.Println("user usecase delete started")
	defer log.Println("user usecase delete done")
	err := input.Validate()
	if err != nil {
		log.Println(err)
		return err
	}

	err = uc.repo.Delete(ctx, input)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
