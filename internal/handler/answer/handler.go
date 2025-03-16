package answer

import (
	"log"
	"net/http"
	"tax-api/internal/entity"
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
	var (
		input    entity.Answer
		output   *entity.Answer
		response entity.Response
		err      error
	)

	log.Println("Upsert answer handle started")
	defer log.Println("Upsert answer handle finished")

	if err = ctx.ShouldBind(&input); err != nil {
		log.Println(err)
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	output, err = h.uc.UpsertAnswer(ctx, input)
	if err != nil {
		log.Println(err)
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}
	response.Message = "ID of created/updated record in Data"
	response.Data = output.ID
	ctx.JSON(http.StatusCreated, response)
	return
}

func (h *Handler) ReadHandle(ctx *gin.Context) {

}

func (h *Handler) DeleteHandle(ctx *gin.Context) {

}
