package handler

import (
	"io"
	"log"
	"net/http"

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
		client := http.Client{} // todo Timeout: time.Second * 3
		req, err := http.NewRequest(http.MethodGet, "http://auth-app/auth/check", nil)
		req.Header.Add("id", ctx.Request.Header.Get("id"))
		req.Header.Add("token", ctx.Request.Header.Get("token"))

		resp, err := client.Do(req)
		if err != nil {
			log.Println(err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Println(err)
			}
		}(resp.Body)
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}
		if resp.StatusCode != http.StatusOK {
			ctx.AbortWithStatusJSON(resp.StatusCode, string(body))
			return
		}

		ctx.Next()
	}
}
