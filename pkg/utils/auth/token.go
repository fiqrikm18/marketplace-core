package auth

import (
	"crypto/rsa"
	"github.com/fiqrikm18/marketplace/core_services/pkg/domain"
	. "github.com/fiqrikm18/marketplace/core_services/pkg/models/API"
	TokenStatus "github.com/fiqrikm18/marketplace/core_services/pkg/types/token/status"
	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"time"
)

const (
	privateKeys = "certs/app.rsa"
	publicKeys  = "certs/app.rsa.pub"
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

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

	accessToken := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), Claims{
		TokenUUID: accessTokenUUID,
		UserID:    data.ID.String(),
		Username:  data.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessExpired,
			Issuer:    "marketplace_core",
		},
	})

	refreshToken := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), Claims{
		TokenUUID: refreshTokenUUID,
		UserID:    data.ID.String(),
		Username:  data.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshExpired,
			Issuer:    "marketplace_core",
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

func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if err := token.Claims.Valid(); err != nil {
			return nil, err
		}

		return verifyKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token.Claims.(*Claims), nil
}

func CheckTokenExpired(expiredTime time.Time) int {
	if time.Now().Before(expiredTime) {
		return TokenStatus.TokenStatusValid
	} else {
		return TokenStatus.TokenStatusExpired
	}
}
