package link

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/rashevskiivv/api/internal/entity"

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

	filter, err = getFilterUS(ctx)
	if err != nil {
		log.Println(err)
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	output, err = h.uc.ReadUS(ctx, filter)
	if err != nil {
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}
	if len(output) == 0 {
		log.Println("Data not found")
		response.Errors = "Data not found"
		ctx.AbortWithStatusJSON(http.StatusNoContent, response)
		return
	}

	response.Data = output
	ctx.JSON(http.StatusOK, response)
	return
}

func getFilterUS(ctx *gin.Context) (filter entity.UserSkillFilter, err error) {
	var val int64
	for k, v := range ctx.Request.URL.Query() {
		switch k {
		case "id_user":
			vals := make([]int64, 0, len(v))
			for _, s := range v {
				val, err = strconv.ParseInt(s, 10, 64)
				if err != nil {
					return filter, err
				}
				vals = append(vals, val)
			}
			filter.UF.ID = vals
		case "id_skill":
			vals := make([]int64, 0, len(v))
			for _, s := range v {
				val, err = strconv.ParseInt(s, 10, 64)
				if err != nil {
					return filter, err
				}
				vals = append(vals, val)
			}
			filter.SF.ID = vals
		case "proficiency_level":
			vals := make([]int64, 0, len(v))
			for _, s := range v {
				val, err = strconv.ParseInt(s, 10, 64)
				if err != nil {
					return filter, err
				}
				vals = append(vals, val)
			}
			filter.ProficiencyLevel = vals
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
		response.Errors = err.Error()
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}
	response.Message = "Deleted"
	ctx.JSON(http.StatusOK, response)
	return
}
