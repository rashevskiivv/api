package link

import (
	usecaseLink "github.com/rashevskiivv/api/internal/usecase/link"
)

type Handler struct {
	uc usecaseLink.UseCaseI
}

func NewHandler(uc usecaseLink.UseCaseI) *Handler {
	return &Handler{uc: uc}
}
