package repositories

import (
	"errors"
	"github.com/fiqrikm18/marketplace/core_services/pkg/domain"
	"gorm.io/gorm"
)

type OauthRepository struct {
	DB *gorm.DB
}

type IOAuthRepository interface {
	Save(data *domain.TokenMeta) error
	Get(uuid string, tokenType int) (*domain.Oauth, error)
	SetStatus(uuid string, status bool) error
}

func NewOauthRepository(db *gorm.DB) *OauthRepository {
	return &OauthRepository{DB: db}
}

func (repo *OauthRepository) Save(data *domain.Oauth) error {
	result := repo.DB.Create(&data)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 || result.RowsAffected == 0 {
		return errors.New("failed inserting data")
	}

	return nil
}

func (repo *OauthRepository) Get(uuid string, tokenType int) (*domain.Oauth, error) {
	return nil, nil
}

func (repo *OauthRepository) SetStatus(uuid string, status bool) error {
	return nil
}
