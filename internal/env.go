package env

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

const (
	AppPort = "APP_PORT"
	DBUrl   = "DB_URL"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("can not find .env file")
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
	dbUrl := os.Getenv(DBUrl)
	if dbUrl == "" {
		return fmt.Sprintf("can not found: %v", DBUrl)
	}
	return dbUrl
}
