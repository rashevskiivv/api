package env

import (
	"log"
	"os"
)

const (
	AppPort = "APP_PORT"
)

func GetEnv() {
	log.Println(os.Getenv(AppPort))
}

func SetEnv() {
	//os.Setenv()
}
