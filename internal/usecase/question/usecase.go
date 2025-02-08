package question

import repositoryQuestion "tax-api/internal/repository/question"

type UseCase struct {
	repo *repositoryQuestion.Repo
}

func NewUseCase(repo *repositoryQuestion.Repo) *UseCase {
	return &UseCase{repo: repo}
}
