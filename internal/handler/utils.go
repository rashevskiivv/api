package handler

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/rashevskiivv/api/internal/entity"
)

const (
	AnswersPath   = "/answers"
	LinksPath     = "/links"
	QuestionsPath = "/questions"
	SkillsPath    = "/skills"
	TestsPath     = "/tests"
	UsersPath     = "/users"
	VacanciesPath = "/vacancies"

	StartPath = "/start"
	EndPath   = "/end"

	TestSkillPath    = "/test_skill"
	UserSkillPath    = "/user_skill"
	SkillVacancyPath = "/skill_vacancy"
)

// NotFound Not found page handler.
func NotFound(c *gin.Context) {
	b := entity.Response{
		Message: "Page not found",
		Errors:  "page not found",
	}
	c.JSON(http.StatusNotFound, b)
}

// HealthCheck Healthcheck page handler.
func HealthCheck(c *gin.Context) {
	b := entity.Response{
		Message: "SERVING",
	}
	c.JSON(http.StatusOK, b)
}

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		client := http.Client{}
		req := http.Request{Method: http.MethodGet, URL: &url.URL{Host: "localhost", Path: "check"}}
		req.Header.Set("id", ctx.Request.Header.Get("id"))
		req.Header.Set("token", ctx.Request.Header.Get("token"))

		resp, err := client.Do(&req)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if resp.StatusCode != http.StatusOK {
			ctx.AbortWithStatus(resp.StatusCode)
			return
		}

		ctx.Next()
	}
}
