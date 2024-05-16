package handler

import (
	"log"
	"net/http"
	"strconv"
	"tax-api/internal/entity"
	"tax-api/internal/repository"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	repo repository.UserRepo
}

func NewUserHandler(repo repository.UserRepo) UserHandler {
	return UserHandler{
		repo: repo,
	}
}

func (h UserHandler) InsertUserHandle(ctx *gin.Context) {
	var (
		user     entity.User
		response entity.Response
	)
	if err := ctx.BindJSON(&user); err != nil {
		log.Println(err)
		response = entity.Response{
			Errors: err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	//todo validate here
	err := h.repo.InsertUser(ctx, user)
	if err != nil {
		log.Println(err)
		response = entity.Response{
			Data:    nil,
			Message: err.Error(),
			Errors:  err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}
	response = entity.Response{
		Message: "Created",
	}
	ctx.JSON(http.StatusCreated, response)
	return
}

func (h UserHandler) ReadUsersHandle(ctx *gin.Context) {
	var (
		filter   entity.Filter
		response entity.Response
		err      error
	)

	if err := ctx.BindJSON(&filter); err != nil {
		log.Println(err)
		response = entity.Response{
			Errors: err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	queryParams := ctx.Request.URL.Query() //todo решить что-то, почему не тело запроса
	filter.Limit, err = strconv.Atoi(queryParams.Get("limit"))
	if err != nil {
		log.Println(err)
		response = entity.Response{
			Errors: err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	filter.Conditions = queryParams //тут и limit

	//todo validate here
	users, err := h.repo.ReadUsers(ctx, filter)
	if err != nil {
		log.Println(err)
		response = entity.Response{
			Data:    nil,
			Message: err.Error(),
			Errors:  err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}
	response = entity.Response{
		Data: users,
	}
	ctx.JSON(http.StatusOK, response)
	return
}
