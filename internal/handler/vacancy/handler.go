package vacancy

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"tax-api/internal/entity"
	usecaseVacancy "tax-api/internal/usecase/vacancy"
	"time"

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
	if err != nil {
		log.Println(err)
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	output, err = h.uc.UpsertVacancy(ctx, input)
	if err != nil {
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

	filter, err = getFilter(ctx)
	if err != nil {
		log.Println(err)
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	answers, err = h.uc.ReadVacancies(ctx, filter)
	if err != nil {
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

func getFilter(ctx *gin.Context) (filter entity.VacancyFilter, err error) {
	var (
		val     int64
		valTime time.Time
	)
	for k, v := range ctx.Request.URL.Query() {
		switch k {
		case "id":
			vals := make([]int64, 0, len(v))
			for _, s := range v {
				val, err = strconv.ParseInt(s, 10, 64)
				if err != nil {
					return filter, err
				}
				vals = append(vals, val)
			}
			filter.ID = vals
		case "title":
			filter.Title = v
		case "grade":
			filter.Grade = v
		case "date":
			vals := make([]time.Time, 0, len(v))
			for _, s := range v {
				valTime, err = time.Parse(time.DateTime, s)
				if err != nil {
					return filter, err
				}
				vals = append(vals, valTime)
			}
			filter.Date = vals
		case "description":
			filter.Description = v
		case "limit":
			if len(v) > 1 {
				err = fmt.Errorf("limit accepts only 1 number")
				return filter, err
			}
			val, err = strconv.ParseInt(v[0], 10, 64)
			if err != nil {
				return filter, err
			}
			filter.Limit = uint(val)
		default:
		}
	}
	return filter, nil
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
	if err != nil {
		log.Println(err)
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = h.uc.DeleteVacancy(ctx, filter)
	if err != nil {
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}
	response.Message = "Deleted"
	ctx.JSON(http.StatusOK, response)
	return
}
