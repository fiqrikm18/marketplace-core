package config

import (
	"fmt"
	"github.com/fiqrikm18/marketplace/core_services/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	DB *gorm.DB
}

func NewConnection() (*DatabaseConfig, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		appConfig.DBHost, appConfig.DBUsername, appConfig.DBPassword, appConfig.DBName, appConfig.DBPort)

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	cnf := DatabaseConfig{
		DB: conn,
	}

	return &cnf, nil
}

func (dbConf *DatabaseConfig) Migrate() {
	err := dbConf.DB.AutoMigrate(&domain.User{}, &domain.Oauth{})
	if err != nil {
		panic(err)
	}
}
