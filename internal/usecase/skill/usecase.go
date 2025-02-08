package skill

import repositorySkill "tax-api/internal/repository/skill"

type UseCase struct {
	repo *repositorySkill.Repo
}

func NewUseCase(repo *repositorySkill.Repo) *UseCase {
	return &UseCase{repo: repo}
}
