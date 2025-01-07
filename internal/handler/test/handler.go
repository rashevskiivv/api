package test

import "github.com/gin-gonic/gin"

type Handler struct {
	uc usecaseTest.UseCaseI
}

func NewTestHandler(uc usecaseTest.UseCaseI) *Handler {
	return &Handler{
		uc: uc,
	}
}

func (h *Handler) UpsertTestHandle(ctx *gin.Context) {

}
func (h *Handler) ReadTestsHandle(ctx *gin.Context) {

}
func (h *Handler) DeleteTestsHandle(ctx *gin.Context) {

}
