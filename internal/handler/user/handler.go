package user

import (
	"errors"
	"io"
	"log"
	"net/http"
	"tax-api/internal/entity"
	usecaseUser "tax-api/internal/usecase/user"

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
		input    entity.User
		output   *entity.User
		response entity.Response
		err      error
	)

	log.Println("Upsert user handle started")
	defer log.Println("Upsert user handle finished")

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

	output, err = h.uc.UpsertUser(ctx, input)
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
		filter   entity.UserFilter
		answers  []entity.User
		response entity.Response
		err      error
	)

	log.Println("Read users handle started")
	defer log.Println("Read users handle finished")

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

	answers, err = h.uc.ReadUsers(ctx, filter)
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
		log.Println(err)
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}
	response.Message = "Deleted"
	ctx.JSON(http.StatusOK, response)
	return
}
