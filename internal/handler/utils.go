package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tax-api/internal/entity"
)

// NotFound Обработчик не найденной страницы
func NotFound(c *gin.Context) {
	b := entity.Response{
		Data:    nil,
		Message: "Page not found",
		Errors:  "page not found",
	}
	c.JSON(http.StatusNotFound, b)
}

// HealthCheck Обработчик хелсчека
func HealthCheck(c *gin.Context) {
	b := entity.Response{
		Data:    nil,
		Message: "SERVING",
		Errors:  "",
	}
	c.JSON(http.StatusOK, b)
}
