package domain

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type TokenMeta struct {
	AccessTokenUUID     string
	RefreshTokenUUID    string
	AccessTokenExpired  int64
	RefreshTokenExpired int64
	AccessTokenString   string
	RefreshTokenString  string
}

type Oauth struct {
	ID                  uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID              uuid.UUID
	User                User
	AccessToken         string
	RefreshToken        string
	AccessTokenExpired  time.Time
	RefreshTokenExpired time.Time
	Expired             bool
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           *time.Time
}

func (u *Oauth) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.NewV4()
	return nil
}
