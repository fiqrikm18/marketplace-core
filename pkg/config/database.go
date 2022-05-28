package config

import (
	"fmt"
	"github.com/fiqrikm18/marketplace/core_services/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type DatabaseConfig struct{}

func (dbConf *DatabaseConfig) NewConnection() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		config.DbHost, config.DbUsername, config.DbPassword, config.DbName, config.DbPort)

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	db = conn
	return nil
}

func (dbConf *DatabaseConfig) Migrate() {
	err := db.AutoMigrate(&models.User{}, &models.Oauth{})
	if err != nil {
		panic(err)
	}
}

func (dbConf *DatabaseConfig) GetConnection() (*gorm.DB, error) {
	if db != nil {
		return db, nil
	}

	err := dbConf.NewConnection()
	if err != nil {
		return nil, err
	}

	return db, nil
}
