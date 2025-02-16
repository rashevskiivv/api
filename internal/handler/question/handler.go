package question

import (
	usecaseQuestion "tax-api/internal/usecase/question"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc usecaseQuestion.UseCaseI
}

func NewHandler(uc usecaseQuestion.UseCaseI) *Handler {
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
