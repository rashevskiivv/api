package skill

import (
	usecaseSkill "tax-api/internal/usecase/skill"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc usecaseSkill.UseCaseI
}

func NewHandler(uc usecaseSkill.UseCaseI) *Handler {
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
