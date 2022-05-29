package models

import (
	"time"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required" valid:"required"`
	Password string `json:"password" binding:"required" valid:"required"`
}

type LoginResponse struct {
	UserId       string    `json:"user_id"`
	AccessToken  string    `json:"access_token"`
	TokenExpired time.Time `json:"token_expired"`
}

type CreateUserRequest struct {
	Username        string `json:"username" valid:"required"`
	Password        string `json:"password" valid:"required"`
	ConfirmPassword string `json:"confirm_password" valid:"required"`
	Email           string `json:"email" valid:"required,email"`
	Phone           string `json:"phone" valid:"required,numeric"`
}

type CreateUserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func (cu *CreateUserRequest) IsPasswordSame() bool {
	if cu.Password != cu.ConfirmPassword {
		return false
	}

	return true
}
