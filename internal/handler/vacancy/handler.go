package vacancy

import (
	usecaseVacancy "tax-api/internal/usecase/vacancy"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc usecaseVacancy.UseCaseI
}

func NewHandler(uc usecaseVacancy.UseCaseI) *Handler {
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
