package test

import (
	usecaseTest "tax-api/internal/usecase/test"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc usecaseTest.UseCaseI
}

func NewHandler(uc usecaseTest.UseCaseI) *Handler {
	return &Handler{
		uc: uc,
	}
}

func (h *Handler) UpsertHandle(ctx *gin.Context) {

}

func (h *Handler) ReadHandle(ctx *gin.Context) {

}

func (h *Handler) DeleteHandle(ctx *gin.Context) {

}
