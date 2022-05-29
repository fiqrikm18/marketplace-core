package config

import (
	"github.com/joho/godotenv"
	"os"
)

type AppConfig struct {
	AppVersion string
	AppPort    string

	DBHost     string
	DBPort     string
	DBUsername string
	DBPassword string
	DBName     string
}

var (
	appConfig *AppConfig
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	configureEnv()
}

func configureEnv() {
	appConfig = &AppConfig{
		AppVersion: os.Getenv("APP_VERSION"),
		AppPort:    os.Getenv("APP_PORT"),

		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUsername: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}
}
