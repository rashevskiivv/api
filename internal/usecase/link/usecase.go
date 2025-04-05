package link

import (
	repositoryLink "github.com/rashevskiivv/api/internal/repository/link"
)

type UseCase struct {
	repo repositoryLink.Repository
}

func NewUseCase(repo repositoryLink.Repository) *UseCase {
	return &UseCase{repo: repo}
}
