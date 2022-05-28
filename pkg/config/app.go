package config

import (
	"errors"
	"os"
)

var (
	config *AppConfig
)

type AppConfig struct {
	// App Configuration
	Version string
	Port    string

	// Postgres DB Configuration
	DbHost     string
	DbPort     string
	DbName     string
	DbUsername string
	DbPassword string
}

func (conf *AppConfig) SetupConfiguration() {
	conf.Version = os.Getenv("APP_VERSION")
	conf.Port = os.Getenv("APP_PORT")

	conf.DbHost = os.Getenv("DB_HOST")
	conf.DbPort = os.Getenv("DB_PORT")
	conf.DbName = os.Getenv("DB_NAME")
	conf.DbUsername = os.Getenv("DB_USERNAME")
	conf.DbPassword = os.Getenv("DB_PASSWORD")

	config = conf
}

func (conf *AppConfig) GetConfiguration() (*AppConfig, error) {
	if config == nil {
		return nil, errors.New("configuration undefined")
	}
	return config, nil
}
