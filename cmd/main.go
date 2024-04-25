package main

import (
	"fmt"
	"log"
	"tax-api/internal"
	"tax-api/internal/handler"

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

	// Running
	err = router.Run(fmt.Sprintf("localhost:%v", appPort))
	if err != nil {
		log.Fatalf("got error while running: %v", err)
	}
}
