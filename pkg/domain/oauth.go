package domain

import (
	TokenStatus "github.com/fiqrikm18/marketplace/core_services/pkg/types/token/status"
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
	AccessTokenUUID     string
	RefreshTokenUUID    string
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

func (u *Oauth) CheckExpiredTime() int {
	if u.Expired {
		return TokenStatus.TOKEN_STATUS_EXPIRED
	}

	remain := time.Unix(u.AccessTokenExpired.Unix(), 0).Sub(time.Now())

	if remain > 0 {
		return TokenStatus.TOKEN_STATUS_EXPIRED
	}

	return TokenStatus.TOKEN_STATUS_EXPIRED
}
