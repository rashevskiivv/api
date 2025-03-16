package question

import (
	"errors"
	"io"
	"log"
	"net/http"
	"tax-api/internal/entity"
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
	var (
		input    entity.Question
		output   *entity.Question
		response entity.Response
		err      error
	)

	log.Println("Upsert question handle started")
	defer log.Println("Upsert question handle finished")

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

	output, err = h.uc.UpsertQuestion(ctx, input)
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
		filter   entity.QuestionFilter
		answers  []entity.Question
		response entity.Response
		err      error
	)

	log.Println("Read questions handle started")
	defer log.Println("Read questions handle finished")

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

	answers, err = h.uc.ReadQuestions(ctx, filter)
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
		filter   entity.QuestionFilter
		response entity.Response
		err      error
	)

	log.Println("Delete question handle started")
	defer log.Println("Delete question handle finished")

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

	err = h.uc.DeleteQuestion(ctx, filter)
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
