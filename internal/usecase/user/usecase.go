package user

import repositoryUser "tax-api/internal/repository/user"

type UseCase struct {
	repo *repositoryUser.Repo
}

func NewUseCase(repo *repositoryUser.Repo) *UseCase {
	return &UseCase{repo: repo}
}
