package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	configureConts "github.com/fiqrikm18/markerplace_core/pkg/const"
)

var (
	appConfig *AppConfig
	err       error
)

type DbConfig struct {
	DB *gorm.DB
}

func NewDBConnection() (*DbConfig, error) {
	appConfig, err = NewAppConfig(configureConts.ConfigurationFileName)
	if err != nil {
		return nil, err
	}

	dbName := appConfig.DBName
	if appConfig.Env == "test" {
		dbName = fmt.Sprintf("%s_test", appConfig.DBName)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		appConfig.DBHost, appConfig.DBUsername, appConfig.DBPassword, dbName, appConfig.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &DbConfig{
		DB: db,
	}, nil
}
