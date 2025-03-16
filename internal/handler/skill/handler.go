package skill

import (
	"errors"
	"io"
	"log"
	"net/http"
	"tax-api/internal/entity"
	usecaseSkill "tax-api/internal/usecase/skill"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc usecaseSkill.UseCaseI
}

func NewHandler(uc usecaseSkill.UseCaseI) *Handler {
	return &Handler{
		uc: uc,
	}
}

func (h *Handler) UpsertHandle(ctx *gin.Context) {
	var (
		input    entity.Skill
		output   *entity.Skill
		response entity.Response
		err      error
	)

	log.Println("Upsert skill handle started")
	defer log.Println("Upsert skill handle finished")

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

	output, err = h.uc.UpsertSkill(ctx, input)
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
		filter   entity.SkillFilter
		skills   []entity.Skill
		response entity.Response
		err      error
	)

	log.Println("Read skills handle started")
	defer log.Println("Read skills handle finished")

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

	skills, err = h.uc.ReadSkills(ctx, filter)
	if err != nil {
		log.Println(err)
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}
	if len(skills) == 0 {
		log.Println("Data not found")
		response.Errors = "Data not found"
		ctx.AbortWithStatusJSON(http.StatusNotFound, response)
		return
	}

	response.Data = skills
	ctx.JSON(http.StatusOK, response)
	return

}

func (h *Handler) DeleteHandle(ctx *gin.Context) {
	var (
		filter   entity.SkillFilter
		response entity.Response
		err      error
	)

	log.Println("Delete skill handle started")
	defer log.Println("Delete skill handle finished")

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

	err = h.uc.DeleteSkill(ctx, filter)
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
