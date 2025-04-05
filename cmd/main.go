package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	repositoryUser "github.com/rashevskiivv/api/internal/repository/user"
	usecaseUser "github.com/rashevskiivv/api/internal/usecase/user"

	"github.com/rashevskiivv/api/internal"
	"github.com/rashevskiivv/api/internal/handler"
	handlerAnswer "github.com/rashevskiivv/api/internal/handler/answer"
	handlerLink "github.com/rashevskiivv/api/internal/handler/link"
	handlerQuestion "github.com/rashevskiivv/api/internal/handler/question"
	handlerSkill "github.com/rashevskiivv/api/internal/handler/skill"
	handlerTest "github.com/rashevskiivv/api/internal/handler/test"
	handlerUser "github.com/rashevskiivv/api/internal/handler/user"
	handlerVacancy "github.com/rashevskiivv/api/internal/handler/vacancy"

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
	handlers := createHandlers(pg)
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

func createHandlers(pg *repository.Postgres) []interface{} {
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
	testUC := usecaseTest.NewUseCase(testRepo, questionRepo, answerRepo)
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

	return []interface{}{testHandler, linkHandler, questionHandler, answerHandler, vacancyHandler, userHandler, skillHandler}
}

func registerHandlers(router *gin.Engine, handlers []interface{}) *gin.Engine {
	// Routing
	router.Use(handler.TokenAuthMiddleware())
	router.NoRoute(handler.NotFound)
	router.GET("/_hc", handler.HealthCheck)

	for _, handlerI := range handlers {
		switch h := handlerI.(type) {
		case *handlerAnswer.Handler:
			router.POST(handler.AnswersPath, h.UpsertHandle)
			router.GET(handler.AnswersPath, h.ReadHandle)
			router.DELETE(handler.AnswersPath, h.DeleteHandle)
			log.Println("answers handler registered")
		case *handlerLink.Handler:
			router.POST(handler.LinksPath+handler.TestSkillPath, h.UpsertTSHandle)
			router.DELETE(handler.LinksPath+handler.TestSkillPath, h.DeleteTSHandle)

			router.POST(handler.LinksPath+handler.UserSkillPath, h.UpsertUSHandle)
			router.GET(handler.LinksPath+handler.UserSkillPath, h.ReadUSHandle)
			router.DELETE(handler.LinksPath+handler.UserSkillPath, h.DeleteUSHandle)

			router.POST(handler.LinksPath+handler.SkillVacancyPath, h.UpsertSVHandle)
			router.GET(handler.LinksPath+handler.SkillVacancyPath, h.ReadSVHandle)
			router.DELETE(handler.LinksPath+handler.SkillVacancyPath, h.DeleteSVHandle)
			log.Println("links handler registered")
		case *handlerQuestion.Handler:
			router.POST(handler.QuestionsPath, h.UpsertHandle)
			router.GET(handler.QuestionsPath, h.ReadHandle)
			router.DELETE(handler.QuestionsPath, h.DeleteHandle)
			log.Println("questions handler registered")
		case *handlerSkill.Handler:
			router.POST(handler.SkillsPath, h.UpsertHandle)
			router.GET(handler.SkillsPath, h.ReadHandle)
			router.DELETE(handler.SkillsPath, h.DeleteHandle)
			log.Println("skills handler registered")
		case *handlerTest.Handler:
			router.POST(handler.TestsPath, h.UpsertHandle)
			router.GET(handler.TestsPath, h.ReadHandle)
			router.DELETE(handler.TestsPath, h.DeleteHandle)
			router.POST(handler.TestsPath+handler.StartPath, h.StartHandle)
			router.POST(handler.TestsPath+handler.EndPath, h.EndHandle)
			log.Println("tests handler registered")
		case *handlerUser.Handler:
			router.POST(handler.UsersPath, h.UpsertHandle)
			router.GET(handler.UsersPath, h.ReadHandle)
			router.DELETE(handler.UsersPath, h.DeleteHandle)
			log.Println("users handler registered")
		case *handlerVacancy.Handler:
			router.POST(handler.VacanciesPath, h.UpsertHandle)
			router.GET(handler.VacanciesPath, h.ReadHandle)
			router.DELETE(handler.VacanciesPath, h.DeleteHandle)
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
