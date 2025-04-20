package user

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
	repositoryUser "github.com/rashevskiivv/api/internal/repository/user"
)

type UseCase struct {
	client *client.Client
	repo   repositoryUser.Repository
}

func NewUseCase(repo repositoryUser.Repository) *UseCase {
	return &UseCase{client: client.NewClient(), repo: repo}
}

func (uc *UseCase) CloseIdleConnections() {
	uc.client.Client.CloseIdleConnections()
}

func (uc *UseCase) UpsertUser(ctx context.Context, input entity.UserAuthInput) (*entity.User, error) {
	var output entity.User
	if input.WhichRequest != entity.AppAuth {
		authAppURL, err := env.GetAuthAppURL()
		if err != nil {
			log.Println(err)
			return nil, err
		}
		out, err := json.Marshal(input.User)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		req := client.NewRequest(http.MethodPost, authAppURL+entity.PathUsers, bytes.NewBuffer(out))
		if req == nil {
			log.Println(err)
			return nil, err
		}

		headers := make(map[string]string, 3)
		headers["id"] = *input.User.ID
		headers["token"] = input.Token
		headers["Origin"] = entity.AppAPI
		req.AddAuthHeaders(headers)

		resp, err := uc.client.Do(req)
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

		body, err := io.ReadAll(resp.Body)
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
		output = entity.User{ID: &id}
	}

	upsertInput := entity.User{
		Name:  input.User.Name,
		Email: input.User.Email,
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

func (uc *UseCase) ReadUsers(ctx context.Context, input entity.UserFilter) ([]entity.User, error) {
	err := input.Validate()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	//output, err := uc.client.Read(ctx, input)
	//if err != nil {
	return nil, err
	//}

	//return output, nil
}

func (uc *UseCase) DeleteUser(ctx context.Context, input entity.UserFilter) error {
	err := input.Validate()
	if err != nil {
		log.Println(err)
		return err
	}
	return err

	//return uc.client.Delete(ctx, input)
}
