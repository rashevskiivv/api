package vacancy

import (
	"errors"
	"io"
	"log"
	"net/http"
	"tax-api/internal/entity"
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
	var (
		input    entity.Vacancy
		output   *entity.Vacancy
		response entity.Response
		err      error
	)

	log.Println("Upsert vacancy handle started")
	defer log.Println("Upsert vacancy handle finished")

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

	output, err = h.uc.UpsertVacancy(ctx, input)
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
		filter   entity.VacancyFilter
		answers  []entity.Vacancy
		response entity.Response
		err      error
	)

	log.Println("Read vacancies handle started")
	defer log.Println("Read vacancies handle finished")

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

	answers, err = h.uc.ReadVacancies(ctx, filter)
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
		filter   entity.VacancyFilter
		response entity.Response
		err      error
	)

	log.Println("Delete vacancy handle started")
	defer log.Println("Delete vacancy handle finished")

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

	err = h.uc.DeleteVacancy(ctx, filter)
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
