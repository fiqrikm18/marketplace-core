package auth

import (
	"crypto/rsa"
	"github.com/fiqrikm18/marketplace/core_services/pkg/domain"
	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"time"
)

const (
	privateKeys = "certs/cert"
	publicKeys  = "certs/cert.pub"
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

type Claims struct {
	TokenUUID string
	UserID    string
	Username  string
}

func init() {
	signBytes, err := ioutil.ReadFile(privateKeys)
	if err != nil {
		panic(err)
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		panic(err)
	}

	verifyBytes, err := ioutil.ReadFile(publicKeys)
	if err != nil {
		panic(err)
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		panic(err)
	}
}

func GenerateToken(data *domain.User) (*domain.TokenMeta, error) {
	accessTokenUUID := uuid.NewV4().String()
	accessExpired := time.Now().Add(time.Hour * 24 * 30).Unix()

	refreshTokenUUID := uuid.NewV4().String()
	refreshExpired := time.Now().Add(time.Hour * 24 * 37).Unix()

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"iss": "marketplace_core",
		"exp": accessExpired,
		"data": Claims{
			TokenUUID: accessTokenUUID,
			UserID:    data.ID.String(),
			Username:  data.Username,
		},
	})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"iss": "marketplace_core",
		"exp": refreshExpired,
		"data": Claims{
			TokenUUID: refreshTokenUUID,
			UserID:    data.ID.String(),
			Username:  data.Username,
		},
	})

	accessTokenString, err := accessToken.SignedString(signKey)
	if err != nil {
		return nil, err
	}

	refreshTokenString, err := refreshToken.SignedString(signKey)
	if err != nil {
		return nil, err
	}

	return &domain.TokenMeta{
		AccessTokenUUID:     accessTokenUUID,
		AccessTokenString:   accessTokenString,
		AccessTokenExpired:  accessExpired,
		RefreshTokenUUID:    refreshTokenUUID,
		RefreshTokenString:  refreshTokenString,
		RefreshTokenExpired: refreshExpired,
	}, nil
}
