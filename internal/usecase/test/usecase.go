package test

import repositoryTest "tax-api/internal/repository/test"

type UseCase struct {
	repo *repositoryTest.Repo
}

func NewUseCase(repo *repositoryTest.Repo) *UseCase {
	return &UseCase{repo: repo}
}
