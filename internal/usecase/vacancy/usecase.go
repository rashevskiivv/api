package vacancy

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"

	env "github.com/rashevskiivv/api/internal"
	"github.com/rashevskiivv/api/internal/client"
	"github.com/rashevskiivv/api/internal/entity"
	repositoryVacancy "github.com/rashevskiivv/api/internal/repository/vacancy"
)

type UseCase struct {
	client *client.Client
	repo   repositoryVacancy.Repository
}

func NewUseCase(repo repositoryVacancy.Repository) *UseCase {
	return &UseCase{client: client.NewClient(), repo: repo}
}

func (uc *UseCase) CloseIdleConnections() {
	uc.client.Client.CloseIdleConnections()
}

func (uc *UseCase) UpsertVacancy(ctx context.Context, input entity.VacancyInput) (*entity.Vacancy, error) {
	var (
		req    *client.Request
		resp   *http.Response
		body   []byte
		output entity.Vacancy
	)
	err := input.Validate()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if input.WhichRequest != entity.AppRecommendations {
		req, err = buildReq(input)
		if err != nil {
			return nil, err
		}
		resp, err = uc.client.Do(req)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		defer func(Body io.ReadCloser) {
			err = Body.Close()
			if err != nil {
				log.Println(err)
			}
		}(resp.Body)

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		response := entity.Response{}
		err = json.Unmarshal(body, &response)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		idf := response.Data.(float64)
		id := int64(math.Round(idf))
		output = entity.Vacancy{ID: &id}
	}

	err = input.Vacancy.Validate()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	upsertInput := entity.Vacancy{
		Title:       input.Vacancy.Title,
		Grade:       input.Vacancy.Grade,
		Date:        input.Vacancy.Date,
		Description: input.Vacancy.Description,
	}
	upsertOutput, err := uc.repo.Upsert(ctx, upsertInput)
	if err != nil {
		return nil, err
	}
	if *upsertOutput.ID != *output.ID {
		log.Println(fmt.Errorf("ids are not the same: id from auth %v, id from api %v", *output.ID, *upsertOutput.ID))
		return nil, fmt.Errorf("ids are not the same at auth and api")
	}

	return &output, nil
}

func buildReq(input entity.VacancyInput) (*client.Request, error) {
	appURL, err := env.GetRecommendationsAppURL()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	out, err := json.Marshal(input.Vacancy)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req := client.NewRequest(http.MethodPost, appURL+entity.PathVacancies, bytes.NewBuffer(out))
	if req == nil {
		log.Println(err)
		return nil, err
	}

	headers := make(map[string]string, 3)
	headers["id"] = input.ID
	headers["token"] = input.Token
	headers["Origin"] = entity.AppAPI
	req.AddAuthHeaders(headers)

	return req, nil
}

func (uc *UseCase) ReadVacancies(ctx context.Context, input entity.VacancyFilter) ([]entity.Vacancy, error) {
	err := input.Validate()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	output, err := uc.repo.Read(ctx, input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (uc *UseCase) DeleteVacancy(ctx context.Context, input entity.VacancyFilter) error {
	err := input.Validate()
	if err != nil {
		log.Println(err)
		return err
	}

	return uc.repo.Delete(ctx, input)
}
