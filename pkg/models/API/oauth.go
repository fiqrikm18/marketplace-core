package API

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	TokenUUID string
	UserID    string
	Username  string

	jwt.StandardClaims
}
