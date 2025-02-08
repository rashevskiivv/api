package main

import (
	"context"
	"fmt"
	"log"

	"tax-api/internal"
	"tax-api/internal/handler"
	handlerAnswer "tax-api/internal/handler/answer"
	handlerQuestion "tax-api/internal/handler/question"
	handlerSkill "tax-api/internal/handler/skill"
	handlerTest "tax-api/internal/handler/test"
	handlerVacancy "tax-api/internal/handler/vacancy"
	"tax-api/internal/repository"
	repositoryAnswer "tax-api/internal/repository/answer"
	repositoryQuestion "tax-api/internal/repository/question"
	repositorySkill "tax-api/internal/repository/skill"
	repositoryTest "tax-api/internal/repository/test"
	repositoryVacancy "tax-api/internal/repository/vacancy"
	usecaseAnswer "tax-api/internal/usecase/answer"
	usecaseQuestion "tax-api/internal/usecase/question"
	usecaseSkill "tax-api/internal/usecase/skill"
	usecaseTest "tax-api/internal/usecase/test"
	usecaseVacancy "tax-api/internal/usecase/vacancy"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	pg, err := getPGInstance()
	if err != nil {
		log.Fatal(err)
	}

	// Routing
	handlers := createHandlers(pg)
	router = registerHandlers(router, handlers)

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

func createHandlers(pg *repository.Postgres) []interface{} {
	// Repo
	answerRepo := repositoryAnswer.NewRepo(pg)
	questionRepo := repositoryQuestion.NewRepo(pg)
	skillRepo := repositorySkill.NewRepo(pg)
	testRepo := repositoryTest.NewRepo(pg)
	userRepo := repositorySkill.NewRepo(pg)
	vacancyRepo := repositoryVacancy.NewRepo(pg)

	// UseCase
	answerUC := usecaseAnswer.NewUseCase(answerRepo)
	questionUC := usecaseQuestion.NewUseCase(questionRepo)
	skillUC := usecaseSkill.NewUseCase(skillRepo)
	testUC := usecaseTest.NewUseCase(testRepo)
	userUC := usecaseSkill.NewUseCase(skillRepo)
	vacancyUC := usecaseVacancy.NewUseCase(vacancyRepo)

	// Handler
	answerHandler := handlerAnswer.NewHandler(answerUC)
	questionHandler := handlerQuestion.NewHandler(questionUC)
	skillHandler := handlerSkill.NewHandler(skillUC)
	testHandler := handlerTest.NewHandler(testUC)
	userHandler := handlerUser.NewHandler(skillUC)
	vacancyHandler := handlerVacancy.NewHandler(vacancyUC)
	//

	log.Println("handlers created")

	return []interface{}{testHandler, questionHandler, answerHandler, vacancyHandler, skillHandler}
}

func registerHandlers(router *gin.Engine, handlers []interface{}) *gin.Engine {
	// Routing
	router.NoRoute(handler.NotFound)
	router.GET("/_hc", handler.HealthCheck)

	for _, handlerI := range handlers {
		switch h := handlerI.(type) {
		case *handlerAnswer.Handler:
			router.POST(handler.AnswersPath, h.UpsertHandle)
			router.GET(handler.AnswersPath, h.ReadHandle)
			router.DELETE(handler.AnswersPath, h.DeleteHandle)
		case *handlerQuestion.Handler:
			router.POST(handler.QuestionsPath, h.UpsertHandle)
			router.GET(handler.QuestionsPath, h.ReadHandle)
			router.DELETE(handler.QuestionsPath, h.DeleteHandle)
		case *handlerSkill.Handler:
			router.POST(handler.SkillPath, h.UpsertHandle)
			router.GET(handler.SkillPath, h.ReadHandle)
			router.DELETE(handler.SkillPath, h.DeleteHandle)
		case *handlerTest.Handler:
			router.POST(handler.TestsPath, h.UpsertHandle)
			router.GET(handler.TestsPath, h.ReadHandle)
			router.DELETE(handler.TestsPath, h.DeleteHandle)
		case *handlerUser.Handler:
			router.POST(handler.UsersPath, h.UpsertHandle)
			router.GET(handler.UsersPath, h.ReadHandle)
			router.DELETE(handler.UsersPath, h.DeleteHandle)
		case *handlerVacancy.Handler:
			router.POST(handler.VacancyPath, h.UpsertHandle)
			router.GET(handler.VacancyPath, h.ReadHandle)
			router.DELETE(handler.VacancyPath, h.DeleteHandle)
		default:
			// todo log
		}
	}

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
