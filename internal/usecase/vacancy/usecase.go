package vacancy

import repositoryVacancy "tax-api/internal/repository/vacancy"

type UseCase struct {
	repo *repositoryVacancy.Repo
}

func NewUseCase(repo *repositoryVacancy.Repo) *UseCase {
	return &UseCase{repo: repo}
}
