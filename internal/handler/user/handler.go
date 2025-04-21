package user

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/rashevskiivv/api/internal/entity"
	usecaseUser "github.com/rashevskiivv/api/internal/usecase/user"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc usecaseUser.UseCaseI
}

func NewHandler(uc usecaseUser.UseCaseI) *Handler {
	return &Handler{
		uc: uc,
	}
}

func (h *Handler) UpsertHandle(ctx *gin.Context) {
	var (
		input    entity.UserAuthInput
		output   *entity.User
		response entity.Response
		err      error
	)

	log.Println("Upsert user handle started")
	defer log.Println("Upsert user handle finished")

	err = ctx.ShouldBind(&input)
	if err != nil {
		log.Println(err)
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	id := ctx.Request.Header.Get("id")
	appSource := ctx.Request.Header.Get("Origin")
	token := ctx.Request.Header.Get("token")
	input.ID = id
	input.WhichRequest = appSource
	input.Token = token
	output, err = h.uc.UpsertUser(ctx, input)
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
		filter   entity.UserFilter
		answers  []entity.User
		response entity.Response
		err      error
	)

	log.Println("Read users handle started")
	defer log.Println("Read users handle finished")

	filter, err = getFilter(ctx)
	if err != nil {
		log.Println(err)
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	answers, err = h.uc.ReadUsers(ctx, filter)
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

func getFilter(ctx *gin.Context) (filter entity.UserFilter, err error) {
	var val int64
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
		case "name":
			filter.Name = v
		case "email":
			filter.Email = v
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
		filter   entity.UserFilter
		response entity.Response
		err      error
	)

	log.Println("Delete user handle started")
	defer log.Println("Delete user handle finished")

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

	err = h.uc.DeleteUser(ctx, filter)
	if err != nil {
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}
	response.Message = "Deleted"
	ctx.JSON(http.StatusOK, response)
	return
}
