package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/rashevskiivv/api/internal"
	"github.com/rashevskiivv/api/internal/entity"
	"github.com/rashevskiivv/api/internal/handler"
	handlerAnswer "github.com/rashevskiivv/api/internal/handler/answer"
	handlerLink "github.com/rashevskiivv/api/internal/handler/link"
	handlerQuestion "github.com/rashevskiivv/api/internal/handler/question"
	handlerSkill "github.com/rashevskiivv/api/internal/handler/skill"
	handlerTest "github.com/rashevskiivv/api/internal/handler/test"
	handlerUser "github.com/rashevskiivv/api/internal/handler/user"
	handlerVacancy "github.com/rashevskiivv/api/internal/handler/vacancy"
	repositoryUser "github.com/rashevskiivv/api/internal/repository/user"

	"github.com/rashevskiivv/api/internal/repository"

	repositoryAnswer "github.com/rashevskiivv/api/internal/repository/answer"
	repositoryLink "github.com/rashevskiivv/api/internal/repository/link"
	repositoryQuestion "github.com/rashevskiivv/api/internal/repository/question"
	repositorySkill "github.com/rashevskiivv/api/internal/repository/skill"
	repositoryTest "github.com/rashevskiivv/api/internal/repository/test"
	repositoryVacancy "github.com/rashevskiivv/api/internal/repository/vacancy"

	usecaseAnswer "github.com/rashevskiivv/api/internal/usecase/answer"
	usecaseLink "github.com/rashevskiivv/api/internal/usecase/link"
	usecaseQuestion "github.com/rashevskiivv/api/internal/usecase/question"
	usecaseSkill "github.com/rashevskiivv/api/internal/usecase/skill"
	usecaseTest "github.com/rashevskiivv/api/internal/usecase/test"
	usecaseUser "github.com/rashevskiivv/api/internal/usecase/user"
	usecaseVacancy "github.com/rashevskiivv/api/internal/usecase/vacancy"

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
	handlers, userUC, vacancyUC := createHandlersAndUCWithClient(pg)
	defer userUC.CloseIdleConnections()
	defer vacancyUC.CloseIdleConnections()
	router = registerHandlers(router, handlers)

	appPort, err := env.GetAppPortEnv()
	if err != nil {
		log.Fatal(err)
	}

	// Running
	err = router.Run(fmt.Sprintf(":%v", strconv.Itoa(appPort)))
	if err != nil {
		log.Fatalf("got error while running: %v", err)
	}
}

func createHandlersAndUCWithClient(pg *repository.Postgres) ([]interface{}, usecaseUser.UseCaseI, usecaseVacancy.UseCaseI) {
	// Repo
	answerRepo := repositoryAnswer.NewRepo(pg)
	linkRepo := repositoryLink.NewRepo(pg)
	questionRepo := repositoryQuestion.NewRepo(pg)
	skillRepo := repositorySkill.NewRepo(pg)
	testRepo := repositoryTest.NewRepo(pg)
	userRepo := repositoryUser.NewRepo(pg)
	vacancyRepo := repositoryVacancy.NewRepo(pg)
	log.Println("repositories created")

	// UseCase
	answerUC := usecaseAnswer.NewUseCase(answerRepo)
	linkUC := usecaseLink.NewUseCase(linkRepo)
	questionUC := usecaseQuestion.NewUseCase(questionRepo)
	skillUC := usecaseSkill.NewUseCase(skillRepo)
	testUC := usecaseTest.NewUseCase(testRepo, questionRepo, answerRepo, linkRepo)
	userUC := usecaseUser.NewUseCase(userRepo)
	vacancyUC := usecaseVacancy.NewUseCase(vacancyRepo)
	log.Println("use cases created")

	// Handler
	answerHandler := handlerAnswer.NewHandler(answerUC)
	linkHandler := handlerLink.NewHandler(linkUC)
	questionHandler := handlerQuestion.NewHandler(questionUC)
	skillHandler := handlerSkill.NewHandler(skillUC)
	testHandler := handlerTest.NewHandler(testUC)
	userHandler := handlerUser.NewHandler(userUC)
	vacancyHandler := handlerVacancy.NewHandler(vacancyUC)
	log.Println("handlers created")

	return []interface{}{testHandler, linkHandler, questionHandler, answerHandler, vacancyHandler, userHandler, skillHandler}, userUC, vacancyUC
}

func registerHandlers(router *gin.Engine, handlers []interface{}) *gin.Engine {
	// Routing
	//router.Use(handler.TokenAuthMiddleware()) // todo uncomment
	router.NoRoute(handler.NotFound)
	router.GET("/_hc", handler.HealthCheck)

	for _, handlerI := range handlers {
		switch h := handlerI.(type) {
		case *handlerAnswer.Handler:
			router.POST(entity.PathAnswers, h.UpsertHandle)
			router.GET(entity.PathAnswers, h.ReadHandle)
			router.DELETE(entity.PathAnswers, h.DeleteHandle)
			log.Println("answers handler registered")
		case *handlerLink.Handler:
			router.POST(entity.PathLinks+entity.PathTestSkill, h.UpsertTSHandle)
			router.DELETE(entity.PathLinks+entity.PathTestSkill, h.DeleteTSHandle)

			router.POST(entity.PathLinks+entity.PathUserSkill, h.UpsertUSHandle)
			router.GET(entity.PathLinks+entity.PathUserSkill, h.ReadUSHandle)
			router.DELETE(entity.PathLinks+entity.PathUserSkill, h.DeleteUSHandle)

			router.POST(entity.PathLinks+entity.PathSkillVacancy, h.UpsertSVHandle)
			router.GET(entity.PathLinks+entity.PathSkillVacancy, h.ReadSVHandle)
			router.DELETE(entity.PathLinks+entity.PathSkillVacancy, h.DeleteSVHandle)
			log.Println("links handler registered")
		case *handlerQuestion.Handler:
			router.POST(entity.PathQuestions, h.UpsertHandle)
			router.GET(entity.PathQuestions, h.ReadHandle)
			router.DELETE(entity.PathQuestions, h.DeleteHandle)
			log.Println("questions handler registered")
		case *handlerSkill.Handler:
			router.POST(entity.PathSkills, h.UpsertHandle)
			router.GET(entity.PathSkills, h.ReadHandle)
			router.DELETE(entity.PathSkills, h.DeleteHandle)
			log.Println("skills handler registered")
		case *handlerTest.Handler:
			router.POST(entity.PathTests, h.UpsertHandle)
			router.GET(entity.PathTests, h.ReadHandle)
			router.DELETE(entity.PathTests, h.DeleteHandle)
			router.POST(entity.PathTests+entity.PathStartTest, h.StartHandle)
			router.POST(entity.PathTests+entity.PathEndTest, h.EndHandle)
			log.Println("tests handler registered")
		case *handlerUser.Handler:
			router.POST(entity.PathUsers, h.UpsertHandle)
			router.GET(entity.PathUsers, h.ReadHandle)
			router.DELETE(entity.PathUsers, h.DeleteHandle)
			log.Println("users handler registered")
		case *handlerVacancy.Handler:
			router.POST(entity.PathVacancies, h.UpsertHandle)
			router.GET(entity.PathVacancies, h.ReadHandle)
			router.DELETE(entity.PathVacancies, h.DeleteHandle)
			log.Println("vacancies handler registered")
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
