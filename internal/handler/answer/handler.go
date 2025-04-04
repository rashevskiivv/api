package answer

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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
	if err != nil {
		log.Println(err)
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	output, err = h.uc.UpsertAnswer(ctx, input)
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
		filter   entity.AnswerFilter
		output   []entity.Answer
		response entity.Response
		err      error
	)

	log.Println("Read answers handle started")
	defer log.Println("Read answers handle finished")

	filter, err = getFilter(ctx)
	if err != nil {
		log.Println(err)
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	output, err = h.uc.ReadAnswers(ctx, filter)
	if err != nil {
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}
	if len(output) == 0 {
		log.Println("Data not found")
		response.Errors = "Data not found"
		ctx.AbortWithStatusJSON(http.StatusNotFound, response)
		return
	}

	response.Data = output
	ctx.JSON(http.StatusOK, response)
	return
}

func getFilter(ctx *gin.Context) (filter entity.AnswerFilter, err error) {
	var (
		val     int64
		valBool bool
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
		case "answer":
			filter.Answer = v
		case "is_right":
			vals := make([]bool, 0, len(v))
			for _, s := range v {
				valBool, err = strconv.ParseBool(s)
				if err != nil {
					return filter, err
				}
				vals = append(vals, valBool)
			}
			filter.IsRight = vals
		case "id_question":
			vals := make([]int64, 0, len(v))
			for _, s := range v {
				val, err = strconv.ParseInt(s, 10, 64)
				if err != nil {
					return filter, err
				}
				vals = append(vals, val)
			}
			filter.IDQuestion = vals
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
		filter   entity.AnswerFilter
		response entity.Response
		err      error
	)

	log.Println("Delete answer handle started")
	defer log.Println("Delete answer handle finished")

	err = ctx.ShouldBind(&filter)
	if err != nil {
		log.Println(err)
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = h.uc.DeleteAnswer(ctx, filter)
	if err != nil {
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}
	response.Message = "Deleted"
	ctx.JSON(http.StatusOK, response)
	return
}
