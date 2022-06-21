package config

import (
	"github.com/spf13/viper"
	"os"
)

type AppConfig struct {
	BaseUrl string
	Port    string
	Env     string
	Version string
	Name    string

	DBHost     string
	DBPort     string
	DBUsername string
	DBPassword string
	DBName     string
}

func NewAppConfig(configName string) (*AppConfig, error) {
	viper.SetConfigName(configName)
	viper.SetConfigType("json")
	viper.AddConfigPath(os.Getenv("MARKETPLACE_CORE_CONFIG"))
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		return nil, err
	}

	return &AppConfig{
		BaseUrl:    viper.GetString("base_url"),
		Port:       viper.GetString("app_port"),
		Env:        viper.GetString("app_environment"),
		Version:    viper.GetString("app_version"),
		Name:       viper.GetString("app_name"),
		DBHost:     viper.GetString("db_host"),
		DBPort:     viper.GetString("db_port"),
		DBUsername: viper.GetString("db_username"),
		DBPassword: viper.GetString("db_password"),
		DBName:     viper.GetString("db_name"),
	}, nil
}
