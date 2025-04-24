package skill

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/rashevskiivv/api/internal/entity"
	usecaseSkill "github.com/rashevskiivv/api/internal/usecase/skill"

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
	if err != nil {
		log.Println(err)
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	output, err = h.uc.UpsertSkill(ctx, input)
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
		filter   entity.SkillFilter
		skills   []entity.Skill
		response entity.Response
		err      error
	)

	log.Println("Read skills handle started")
	defer log.Println("Read skills handle finished")

	filter, err = getFilter(ctx)
	if err != nil {
		log.Println(err)
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	skills, err = h.uc.ReadSkills(ctx, filter)
	if err != nil {
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}
	if len(skills) == 0 {
		log.Println("Data not found")
		response.Errors = "Data not found"
		ctx.AbortWithStatusJSON(http.StatusNoContent, response)
		return
	}

	response.Data = skills
	ctx.JSON(http.StatusOK, response)
	return
}

func getFilter(ctx *gin.Context) (filter entity.SkillFilter, err error) {
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
		case "title":
			filter.Title = v
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
		filter   entity.SkillFilter
		response entity.Response
		err      error
	)

	log.Println("Delete skill handle started")
	defer log.Println("Delete skill handle finished")

	err = ctx.ShouldBind(&filter)
	if err != nil {
		log.Println(err)
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = h.uc.DeleteSkill(ctx, filter)
	if err != nil {
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}
	response.Message = "Deleted"
	ctx.JSON(http.StatusOK, response)
	return
}
