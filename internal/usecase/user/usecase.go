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

	output, err := uc.repo.Upsert(ctx, input)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("user usecase upsert done")

	return output, nil
}

func (uc *UseCase) ReadUsers(ctx context.Context, input entity.UserFilter) ([]entity.User, error) {
	log.Println("user usecase read started")
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
	log.Println("user usecase read done")

	return output, nil
}

func (uc *UseCase) DeleteUser(ctx context.Context, input entity.UserFilter) error {
	log.Println("user usecase delete started")
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
	log.Println("user usecase delete done")

	return nil
}
