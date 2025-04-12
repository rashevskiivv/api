package env

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	envAppPort          = "APP_PORT"
	envPostgresDriver   = "POSTGRES_DRIVER"
	envPostgresUser     = "POSTGRES_USER"
	envPostgresPassword = "POSTGRES_PASSWORD"
	envPostgresHost     = "POSTGRES_HOST"
	envPostgresPort     = "POSTGRES_PORT"
	envPostgresDB       = "POSTGRES_DB"
	envAuthAppURL       = "AUTH_APP_URL"
)

func init() {
	err := godotenv.Load("deployment/.env")
	if err != nil {
		log.Fatal("can not find .env file: ", err)
	}
	log.Println("env loaded")
}

func GetAuthAppURL() (string, error) {
	authEnv := os.Getenv(envAuthAppURL)
	if authEnv == "" {
		return "", fmt.Errorf("can not found: %v", envAuthAppURL)
	}
	return authEnv, nil
}

func GetAppPortEnv() (int, error) {
	portStr := os.Getenv(envAppPort)
	if portStr == "" {
		return 0, fmt.Errorf("can not found: %v", envAppPort)
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return 0, fmt.Errorf("can not convert to integer: %v", envAppPort)
	}
	log.Println("app port got from env")
	return port, nil
}

func GetDBUrlEnv() (string, error) {
	dbDriver := os.Getenv(envPostgresDriver)
	if dbDriver == "" {
		return "", fmt.Errorf("can not found: %v", envPostgresDriver)
	}
	dbUser := os.Getenv(envPostgresUser)
	if dbUser == "" {
		return "", fmt.Errorf("can not found: %v", envPostgresUser)
	}
	dbPassword := os.Getenv(envPostgresPassword)
	if dbPassword == "" {
		return "", fmt.Errorf("can not found: %v", envPostgresPassword)
	}
	dbHost := os.Getenv(envPostgresHost)
	if dbHost == "" {
		return "", fmt.Errorf("can not found: %v", envPostgresHost)
	}
	dbPort := os.Getenv(envPostgresPort)
	if dbPort == "" {
		return "", fmt.Errorf("can not found: %v", envPostgresPort)
	}
	dbName := os.Getenv(envPostgresDB)
	if dbName == "" {
		return "", fmt.Errorf("can not found: %v", envPostgresDB)
	}
	log.Println("db url got from env")
	return fmt.Sprintf("%v://%v:%v@%v:%v/%v", dbDriver, dbUser, dbPassword, dbHost, dbPort, dbName), nil
}
