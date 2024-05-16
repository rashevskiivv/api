package main

import (
	"context"
	"fmt"
	"log"
	"tax-api/internal"
	"tax-api/internal/handler"
	"tax-api/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Routing
	router.NoRoute(handler.NotFound)
	router.GET("/_hc", handler.HealthCheck)

	appPort, err := env.GetAppPortEnv()
	if err != nil {
		log.Fatal(err)
	}

	router = registerHandlers(router)

	// Running
	err = router.Run(fmt.Sprintf("localhost:%v", appPort))
	if err != nil {
		log.Fatalf("got error while running: %v", err)
	}
}

func registerHandlers(router *gin.Engine) *gin.Engine {
	ctx := context.Background()
	userRepo := repository.NewUserRepo(ctx)
	h := handler.NewUserHandler(userRepo)

	router.POST("users", h.InsertUserHandle)
	router.GET("users", h.ReadUsersHandle)

	return router
}
