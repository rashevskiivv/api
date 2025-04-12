package user

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	env "github.com/rashevskiivv/api/internal"
	"github.com/rashevskiivv/api/internal/client"
	"github.com/rashevskiivv/api/internal/entity"
)

type UseCase struct {
	client *client.Client
}

func NewUseCase() *UseCase {
	return &UseCase{client: client.NewClient()}
}

func (uc *UseCase) CloseIdleConnections() {
	uc.client.Client.CloseIdleConnections()
}

func (uc *UseCase) UpsertUser(ctx context.Context, input entity.UserAuthInput) (*entity.User, error) {
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

	headers := make(map[string]string, 2)
	headers["id"] = *input.User.ID
	headers["token"] = input.Token
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
	id := response.Data.(int64)
	output := entity.User{ID: &id}

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
