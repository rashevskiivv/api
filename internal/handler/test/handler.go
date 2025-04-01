package test

import (
	"log"
	"net/http"
	"tax-api/internal/entity"
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
	var (
		input    entity.Test
		output   *entity.Test
		response entity.Response
		err      error
	)

	log.Println("Upsert test handle started")
	defer log.Println("Upsert test handle finished")

	err = ctx.ShouldBind(&input)
	if err != nil {
		log.Println(err)
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	output, err = h.uc.UpsertTest(ctx, input)
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
		filter   entity.TestFilter
		answers  []entity.Test
		response entity.Response
		err      error
	)

	log.Println("Read tests handle started")
	defer log.Println("Read tests handle finished")

	err = ctx.ShouldBind(&filter)
	if err != nil {
		log.Println(err)
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	answers, err = h.uc.ReadTests(ctx, filter)
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
		filter   entity.TestFilter
		response entity.Response
		err      error
	)

	log.Println("Delete test handle started")
	defer log.Println("Delete test handle finished")

	err = ctx.ShouldBind(&filter)
	if err != nil {
		log.Println(err)
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = h.uc.DeleteTest(ctx, filter)
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

func (h *Handler) StartHandle(ctx *gin.Context) {
	var (
		input    entity.StartTestInput
		output   *entity.StartTestOutput
		response entity.Response
		err      error
	)

	log.Println("Start test handle started")
	defer log.Println("Start test handle finished")

	err = ctx.ShouldBind(&input)
	if err != nil {
		log.Println(err)
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	output, err = h.uc.StartTest(ctx, input)
	if err != nil {
		log.Println(err)
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	response.Data = output
	ctx.JSON(http.StatusOK, response)
	return
}

func (h *Handler) EndHandle(ctx *gin.Context) {
	// todo implement me
}
