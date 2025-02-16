package answer

import (
	usecaseAnswer "tax-api/internal/usecase/answer"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc usecaseAnswer.UseCaseI
}

func NewHandler(uc usecaseAnswer.UseCaseI) *Handler {
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
