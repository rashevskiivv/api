package answer

import (
	"errors"
	"io"
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
	return &Handler{uc: uc}
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

	err = ctx.ShouldBind(&input)
	if errors.Is(err, io.EOF) {
		err = nil
	}
	if err != nil {
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
	var (
		filter   entity.AnswerFilter
		answers  []entity.Answer
		response entity.Response
		err      error
	)

	log.Println("Read answers handle started")
	defer log.Println("Read answers handle finished")

	err = ctx.ShouldBind(&filter)
	if errors.Is(err, io.EOF) {
		err = nil
	}
	if err != nil {
		log.Println(err)
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	answers, err = h.uc.ReadAnswers(ctx, filter)
	if err != nil {
		log.Println(err)
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}
	if len(answers) == 0 {
		log.Println("Data not found")
		response.Errors = "Data not found"
		ctx.AbortWithStatusJSON(http.StatusNotFound, response)
		return
	}

	response.Data = answers
	ctx.JSON(http.StatusOK, response)
	return
}

func (h *Handler) DeleteHandle(ctx *gin.Context) {
	var (
		filter   entity.AnswerFilter
		response entity.Response
		err      error
	)

	log.Println("Delete answer handle started")
	defer log.Println("Delete answer handle finished")

	err = ctx.ShouldBind(&filter)
	if errors.Is(err, io.EOF) {
		err = nil
	}
	if err != nil {
		log.Println(err)
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = h.uc.DeleteAnswer(ctx, filter)
	if err != nil {
		log.Println(err)
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}
	response.Message = "Deleted"
	ctx.JSON(http.StatusOK, response)
	return
}
