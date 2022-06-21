package domain

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	Name        string    `gorm:"not null;index"`
	Username    string    `gorm:"not null;index;unique"`
	Email       string    `gorm:"not null;index;unique"`
	PhoneNumber string    `gorm:"not null"`
	Gender      int
	BirthDay    time.Time
	BirthPlace  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewV4()
	if err != nil {
		return err
	}

	if u.Name == "" {
		return errors.New("name cannot be empty")
	}

	if u.Username == "" {
		return errors.New("username cannot be empty")
	}

	if u.Email == "" {
		return errors.New("email cannot be empty")
	}

	if !govalidator.IsEmail(u.Email) {
		return errors.New("invalid email format")
	}

	if u.PhoneNumber == "" {
		return errors.New("phone number cannot be empty")
	}

	return nil
}
