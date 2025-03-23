package link

import (
	"log"
	"net/http"
	"tax-api/internal/entity"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UpsertSVHandle(ctx *gin.Context) {
	var (
		input    entity.SkillVacancy
		response entity.Response
		err      error
	)

	log.Println("Upsert skill-vacancy handle started")
	defer log.Println("Upsert skill-vacancy handle finished")

	err = ctx.ShouldBind(&input)
	if err != nil {
		log.Println(err)
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = h.uc.UpsertSV(ctx, input)
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

func (h *Handler) ReadSVHandle(ctx *gin.Context) {
	var (
		filter   entity.SkillVacancyFilter
		output   []entity.SkillVacancy
		response entity.Response
		err      error
	)

	log.Println("Read skill-vacancy handle started")
	defer log.Println("Read skill-vacancy handle finished")

	err = ctx.ShouldBind(&filter)
	if err != nil {
		log.Println(err)
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	output, err = h.uc.ReadSV(ctx, filter)
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

func (h *Handler) DeleteSVHandle(ctx *gin.Context) {
	var (
		filter   entity.SkillVacancyFilter
		response entity.Response
		err      error
	)

	log.Println("Delete skill-vacancy handle started")
	defer log.Println("Delete skill-vacancy handle finished")

	err = ctx.ShouldBind(&filter)
	if err != nil {
		log.Println(err)
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = h.uc.DeleteSV(ctx, filter)
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
