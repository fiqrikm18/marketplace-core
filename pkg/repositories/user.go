package repositories

import (
	"errors"
	"github.com/fiqrikm18/marketplace/core_services/pkg/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type IUserRepository interface {
	CreateUser(data *domain.User) (string, error)
	UpdateUser(uuid string, data *domain.User) error
	DeleteUser(uuid string) error
	GetUserByUUID(uuid string) (*domain.User, error)
	GetAllUser() (*[]domain.User, error)
	ValidateUserExist(username, email string) (*domain.User, error)
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) CreateUser(data *domain.User) (string, error) {
	result := repo.db.Create(&data)
	if result.Error != nil {
		return "", result.Error
	}

	if result.RowsAffected < 1 || result.RowsAffected == 0 {
		return "", errors.New("failed inserting data")
	}

	return data.ID.String(), nil
}

func (repo *UserRepository) UpdateUser(uuid string, data *domain.User) error {
	return nil
}
func (repo *UserRepository) DeleteUser(uuid string) error {
	return nil
}

func (repo *UserRepository) GetUserByUUID(uuid string) (*domain.User, error) {
	var user domain.User
	trx := repo.db.Where("id = ?", uuid).First(&user)
	if trx.Error != nil {
		return nil, trx.Error
	}

	if user.ID.String() == "" {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (repo *UserRepository) GetUserByUsername(username string) (*domain.User, error) {
	var user domain.User
	trx := repo.db.Where("username = ?", username).First(&user)
	if trx.Error != nil {
		return nil, trx.Error
	}

	if user.ID.String() == "" {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (repo *UserRepository) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	trx := repo.db.Where("email = ?", email).First(&user)
	if trx.Error != nil {
		return nil, trx.Error
	}

	if user.ID.String() == "" {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (repo *UserRepository) GetAllUser() (*[]domain.User, error) {
	return nil, nil
}
