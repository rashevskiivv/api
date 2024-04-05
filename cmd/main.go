package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"tax-api/internal/handler"
)

func main() {
	router := gin.Default()

	// Middlewares
	router.Use(gin.Recovery())

	// Routing
	router.NoRoute(handler.NotFound)
	router.GET("/_hc", handler.HealthCheck)

	// Running
	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatalf("got error while running: %v", err)
	}
}
