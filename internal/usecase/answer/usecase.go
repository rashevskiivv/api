package answer

import repositoryAnswer "tax-api/internal/repository/answer"

type UseCase struct {
	repo *repositoryAnswer.Repo
}

func NewUseCase(repo *repositoryAnswer.Repo) *UseCase {
	return &UseCase{repo: repo}
}
