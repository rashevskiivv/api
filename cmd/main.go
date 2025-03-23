package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	repositoryUser "tax-api/internal/repository/user"
	usecaseUser "tax-api/internal/usecase/user"

	"tax-api/internal"
	"tax-api/internal/handler"

	handlerAnswer "tax-api/internal/handler/answer"
	handlerLink "tax-api/internal/handler/link"
	handlerQuestion "tax-api/internal/handler/question"
	handlerSkill "tax-api/internal/handler/skill"
	handlerTest "tax-api/internal/handler/test"
	handlerUser "tax-api/internal/handler/user"
	handlerVacancy "tax-api/internal/handler/vacancy"

	"tax-api/internal/repository"

	repositoryAnswer "tax-api/internal/repository/answer"
	repositoryLink "tax-api/internal/repository/link"
	repositoryQuestion "tax-api/internal/repository/question"
	repositorySkill "tax-api/internal/repository/skill"
	repositoryTest "tax-api/internal/repository/test"
	repositoryVacancy "tax-api/internal/repository/vacancy"

	usecaseAnswer "tax-api/internal/usecase/answer"
	usecaseLink "tax-api/internal/usecase/link"
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
	testUC := usecaseTest.NewUseCase(testRepo)
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
			router.GET(handler.LinksPath+handler.TestSkillPath, h.ReadTSHandle)
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
