package env

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	AppPort     = "APP_PORT"
	PostgresUrl = "POSTGRES_URL"
)

func init() {
	if err := godotenv.Load("deployment/.env"); err != nil {
		log.Fatal("can not find .env file: ", err)
	}
}

func GetAppPortEnv() (int, error) {
	portStr := os.Getenv(AppPort)
	if portStr == "" {
		return 0, errors.New(fmt.Sprintf("can not found: %v", AppPort))
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("can not convert to integer: %v", AppPort))
	}
	return port, nil
}

func GetDBUrlEnv() string {
	dbUrl := os.Getenv(PostgresUrl)
	if dbUrl == "" {
		return fmt.Sprintf("can not found: %v", PostgresUrl)
	}
	return dbUrl
}
