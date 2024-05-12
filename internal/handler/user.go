package handler

import (
	"log"
	"net/http"
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
	}
	response = entity.Response{
		Data:    nil,
		Message: "Created",
		Errors:  "",
	}
	ctx.JSON(http.StatusCreated, response)
}
