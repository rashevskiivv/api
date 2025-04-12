package handler

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	env "github.com/rashevskiivv/api/internal"
	"github.com/rashevskiivv/api/internal/client"
	"github.com/rashevskiivv/api/internal/entity"
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
		cl := client.NewClient()
		defer cl.Client.CloseIdleConnections()

		authAppURL, err := env.GetAuthAppURL()
		if err != nil {
			log.Println(err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}
		req := client.NewRequest(http.MethodGet, authAppURL+entity.PathCheck, nil)
		if req == nil {
			log.Println(entity.ErrorCreateRequest)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, entity.ErrorCreateRequest)
			return
		}

		headers := make(map[string]string, 2)
		headers["id"] = ctx.Request.Header.Get("id")
		headers["token"] = ctx.Request.Header.Get("token")
		req.AddAuthHeaders(headers)

		resp, err := cl.Do(req)
		if err != nil {
			log.Println(err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		defer func(Body io.ReadCloser) {
			err = Body.Close()
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
