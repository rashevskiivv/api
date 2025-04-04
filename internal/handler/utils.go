package handler

import (
	"net/http"
	"tax-api/internal/entity"

	"github.com/gin-gonic/gin"
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
