package env

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"tax-api/internal/entity"

	"github.com/joho/godotenv"
)

const (
	AppPort = "APP_PORT"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("can not find .env file")
	}
}

func GetEnv() (*entity.Env, error) {
	portStr := os.Getenv(AppPort)
	if portStr == "" {
		return nil, errors.New(fmt.Sprintf("can not found %v", AppPort))
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("can not found %v", AppPort))
	}
	return &entity.Env{AppPort: port}, nil
}
