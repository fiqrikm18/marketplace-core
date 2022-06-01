package repositories

import (
	"errors"
	"github.com/fiqrikm18/marketplace/core_services/pkg/domain"
	"github.com/fiqrikm18/marketplace/core_services/pkg/types/token/type"
	"gorm.io/gorm"
)

type OauthRepository struct {
	DB *gorm.DB
}

type IOAuthRepository interface {
	Get(uuid string, tokenType int) (*domain.Oauth, error)
	Save(data *domain.Oauth) error
	Revoke(uuid string) error
	RevokeAll(userID string) error
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
	var oauth domain.Oauth
	if tokenType == _type.AccessToken {
		tx := repo.DB.Where("access_token_uuid=?", uuid).Order("created_at desc").First(&oauth)
		if tx.Error != nil {
			return nil, tx.Error
		}
	} else if tokenType == _type.RefreshToken {
		tx := repo.DB.Where("refresh_token_uuid=?", uuid).Order("created_at desc").First(&oauth)
		if tx.Error != nil {
			return nil, tx.Error
		}
	} else {
		return nil, errors.New("invalid token type")
	}

	return &oauth, nil
}

func (repo *OauthRepository) Revoke(uuid string) error {
	tx := repo.DB.Model(&domain.Oauth{}).Where("access_token_uuid = ?", uuid).Update("expired", true)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (repo *OauthRepository) RevokeAll(userID string) error {
	tx := repo.DB.Model(&domain.Oauth{}).Where("user_id = ?", userID).Update("expired", true)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
