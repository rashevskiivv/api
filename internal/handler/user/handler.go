package user

import (
	usecaseUser "tax-api/internal/usecase/user"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc usecaseUser.UseCaseI
}

func NewHandler(uc usecaseUser.UseCaseI) *Handler {
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
