package lib

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type EnvironmentVariables struct {
	AppServerAddress string
	AllowedIps       string
	GoogleClientId   string
}

func ReadEnvironmentVariables() EnvironmentVariables {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file: %v\n", err)
	}

	appAddress := os.Getenv("APP_SERVER_ADDR")
	if appAddress == "" {
		appAddress = ":8080"
	}
	return EnvironmentVariables{
		AppServerAddress: appAddress,
		AllowedIps:       os.Getenv("ALLOWED_IPS"),
		GoogleClientId:   os.Getenv("GOOGLE_SERVER_CLIENT_ID"),
	}
}
