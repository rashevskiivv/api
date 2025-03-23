package link

import (
	"log"
	"net/http"
	"tax-api/internal/entity"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UpsertTSHandle(ctx *gin.Context) {
	var (
		input    entity.TestSkill
		response entity.Response
		err      error
	)

	log.Println("Upsert test-skill handle started")
	defer log.Println("Upsert test-skill handle finished")

	err = ctx.ShouldBind(&input)
	if err != nil {
		log.Println(err)
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = h.uc.UpsertTS(ctx, input)
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

func (h *Handler) ReadTSHandle(ctx *gin.Context) {
	var (
		filter   entity.TestSkillFilter
		output   []entity.TestSkill
		response entity.Response
		err      error
	)

	log.Println("Read test-skill handle started")
	defer log.Println("Read test-skill handle finished")

	err = ctx.ShouldBind(&filter)
	if err != nil {
		log.Println(err)
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	output, err = h.uc.ReadTS(ctx, filter)
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

func (h *Handler) DeleteTSHandle(ctx *gin.Context) {
	var (
		filter   entity.TestSkillFilter
		response entity.Response
		err      error
	)

	log.Println("Delete test-skill handle started")
	defer log.Println("Delete test-skill handle finished")

	err = ctx.ShouldBind(&filter)
	if err != nil {
		log.Println(err)
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = h.uc.DeleteTS(ctx, filter)
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
