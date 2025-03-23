package link

import (
	"log"
	"net/http"
	"tax-api/internal/entity"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UpsertUSHandle(ctx *gin.Context) {
	var (
		input    entity.UserSkill
		response entity.Response
		err      error
	)

	log.Println("Upsert user-skill handle started")
	defer log.Println("Upsert user-skill handle finished")

	err = ctx.ShouldBind(&input)
	if err != nil {
		log.Println(err)
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = h.uc.UpsertUS(ctx, input)
	if err != nil {
		log.Println(err)
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	response.Message = "Link created"
	ctx.JSON(http.StatusCreated, response)
	return
}

func (h *Handler) ReadUSHandle(ctx *gin.Context) {
	var (
		filter   entity.UserSkillFilter
		output   []entity.UserSkill
		response entity.Response
		err      error
	)

	log.Println("Read user-skill handle started")
	defer log.Println("Read user-skill handle finished")

	err = ctx.ShouldBind(&filter)
	if err != nil {
		log.Println(err)
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	output, err = h.uc.ReadUS(ctx, filter)
	if err != nil {
		log.Println(err)
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

func (h *Handler) DeleteUSHandle(ctx *gin.Context) {
	var (
		filter   entity.UserSkillFilter
		response entity.Response
		err      error
	)

	log.Println("Delete user-skill handle started")
	defer log.Println("Delete user-skill handle finished")

	err = ctx.ShouldBind(&filter)
	if err != nil {
		log.Println(err)
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = h.uc.DeleteUS(ctx, filter)
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
