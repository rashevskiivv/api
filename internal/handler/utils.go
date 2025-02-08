package handler

import (
	"net/http"
	"tax-api/internal/entity"

	"github.com/gin-gonic/gin"
)

const (
	AnswersPath   = "/answers"
	QuestionsPath = "/questions"
	SkillPath     = "/skills"
	TestsPath     = "/tests"
	UsersPath     = "/users"
	VacancyPath   = "/vacancies"
)

// NotFound Not found page handler.
func NotFound(c *gin.Context) {
	b := entity.Response{
		Data:    nil,
		Message: "Page not found",
		Errors:  "page not found",
	}
	c.JSON(http.StatusNotFound, b)
}

// HealthCheck Healthcheck page handler.
func HealthCheck(c *gin.Context) {
	b := entity.Response{
		Data:    nil,
		Message: "SERVING",
		Errors:  "",
	}
	c.JSON(http.StatusOK, b)
}
