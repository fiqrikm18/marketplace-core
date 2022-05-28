package API

import "time"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserId       string    `json:"user_id"`
	AccessToken  string    `json:"access_token"`
	TokenExpired time.Time `json:"token_expired"`
}
