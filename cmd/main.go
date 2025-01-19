package main

import (
	"context"
	"fmt"
	"log"

	"tax-api/internal"
	"tax-api/internal/handler"
	"tax-api/internal/repository"
	repositoryAnswer "tax-api/internal/repository/answer"
	repositoryQuestion "tax-api/internal/repository/question"
	repositoryTest "tax-api/internal/repository/test"
	repositoryVacancy "tax-api/internal/repository/vacancy"

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

func createHandlers(pg *repository.Postgres) {
	// Repo
	testRepo := repositoryTest.NewRepo(pg)
	questionRepo := repositoryQuestion.NewRepo(pg)
	answerRepo := repositoryAnswer.NewRepo(pg)
	vacancyRepo := repositoryVacancy.NewRepo(pg)

	// UseCase

	// Handler
	//

	log.Println("handlers created")

	return
}

func registerHandlers(router *gin.Engine) *gin.Engine {
	// Routing
	router.NoRoute(handler.NotFound)
	router.GET("/_hc", handler.HealthCheck)

	return router
}

func getPGInstance() (*repository.Postgres, error) {
	url, err := env.GetDBUrlEnv()
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	pg, err := repository.NewPG(ctx, url)
	if err != nil {
		return nil, err
	}
	return pg, nil
}
