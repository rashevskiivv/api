package link

import (
	repositoryLink "tax-api/internal/repository/link"
)

type UseCase struct {
	repo repositoryLink.Repository
}

func NewUseCase(repo repositoryLink.Repository) *UseCase {
	return &UseCase{repo: repo}
}
