package repositories

import (
	"errors"
	"github.com/fiqrikm18/markerplace_core/pkg/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

type IUserRepository interface {
	Save(user *domain.User) error
	Update(uuid string, user *domain.User) (*domain.User, error)
	Delete(uuid string) error
	DeleteAll() error
	GetUserById(uuid string) (*domain.User, error)
	GetUsers() (*[]domain.User, error)
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (u *UserRepository) Save(user *domain.User) error {
	trx := u.DB.Create(user)
	if err := trx.Error; err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) Update(uuid string, user *domain.User) (*domain.User, error) {
	return nil, nil
}

func (u *UserRepository) Delete(uuid string) error {
	return errors.New("not implemented")
}

func (u *UserRepository) DeleteAll() error {
	return nil
}

func (u *UserRepository) GetUserById(uuid string) (*domain.User, error) {
	var user domain.User
	trx := u.DB.Where("id = ?", uuid).First(&user)
	if err := trx.Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) GetUsers() (*[]domain.User, error) {
	return nil, nil
}
