package auth

import (
	"github.com/asaskevich/govalidator"
	"github.com/fiqrikm18/marketplace/core_services/pkg/config"
	"github.com/fiqrikm18/marketplace/core_services/pkg/domain"
	"github.com/fiqrikm18/marketplace/core_services/pkg/models"
	API2 "github.com/fiqrikm18/marketplace/core_services/pkg/models/API"
	"github.com/fiqrikm18/marketplace/core_services/pkg/repositories"
	_type "github.com/fiqrikm18/marketplace/core_services/pkg/types/token/type"
	"github.com/fiqrikm18/marketplace/core_services/pkg/utils/API"
	"github.com/fiqrikm18/marketplace/core_services/pkg/utils/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"strings"
	"time"
)

var (
	userRepository  *repositories.UserRepository
	oauthRepository *repositories.OauthRepository
)

func init() {
	dbConf, err := config.NewConnection()
	if err != nil {
		panic(err)
	}

	userRepository = repositories.NewUserRepository(dbConf.DB)
	oauthRepository = repositories.NewOauthRepository(dbConf.DB)
}

func Login(ctx *gin.Context) {
	var payload models.LoginRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		API.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	if govalidator.IsNull(payload.Username) {
		API.ErrorResponse(ctx, http.StatusUnprocessableEntity, "username is required")
	}

	if govalidator.IsNull(payload.Password) {
		API.ErrorResponse(ctx, http.StatusUnprocessableEntity, "password is required")
	}

	user, err := userRepository.GetUserByUsername(payload.Username)
	if err != nil {
		API.ErrorResponse(ctx, http.StatusNotFound, "user not registered")
	}

	if govalidator.IsNull(user.Username) {
		API.ErrorResponse(ctx, http.StatusUnauthorized, "invalid username or password")
	}

	err = user.ValidatePassword(payload.Password)
	if err != nil {
		API.ErrorResponse(ctx, http.StatusUnauthorized, "invalid username or password")
	}

	tokenData, err := auth.GenerateToken(user)
	if err != nil {
		API.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	token := domain.Oauth{
		UserID:              user.ID,
		AccessTokenUUID:     tokenData.AccessTokenUUID,
		RefreshTokenUUID:    tokenData.RefreshTokenUUID,
		AccessTokenExpired:  time.Unix(tokenData.AccessTokenExpired, 0),
		RefreshTokenExpired: time.Unix(tokenData.RefreshTokenExpired, 0),
		Expired:             false,
	}

	err = oauthRepository.Save(&token)
	if err != nil {
		API.ErrorResponse(ctx, http.StatusInternalServerError, "internal server error")
	}

	ctx.SetCookie("access_token", tokenData.AccessTokenString, int(tokenData.AccessTokenExpired), "/", ctx.Request.Host, true, true)
	ctx.SetCookie("refresh_token", tokenData.RefreshTokenString, int(tokenData.RefreshTokenExpired), "/", ctx.Request.Host, true, true)

	API.SuccessResponse(ctx, http.StatusOK, "", API2.M{
		"access_token":  tokenData.AccessTokenString,
		"refresh_token": tokenData.RefreshTokenString,
		"expired_at":    tokenData.AccessTokenExpired,
	})
}

func Logout(ctx *gin.Context) {
	tokenString := strings.Split(ctx.Request.Header["Authorization"][0], " ")[1]
	tokenClaims, err := auth.ParseToken(tokenString)
	if err != nil {
		API.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	token, err := oauthRepository.Get(tokenClaims.TokenUUID, _type.ACCESS_TOKEN)
	if err != nil {
		API.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	err = oauthRepository.Revoke(token.AccessTokenUUID)
	if err != nil {
		API.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	API.SuccessResponse(ctx, http.StatusOK, "logout success", nil)
}

func Register(ctx *gin.Context) {
	var payload models.CreateUserRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		API.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	phoneRegex, err := regexp.Compile("\\d+")
	if err != nil {
		API.ErrorResponse(ctx, http.StatusInternalServerError, "internal server error")
	}

	if !phoneRegex.MatchString(payload.Phone) {
		API.ErrorResponse(ctx, http.StatusUnprocessableEntity, "phone must be numeric")
	}

	if govalidator.IsNull(payload.Email) {
		API.ErrorResponse(ctx, http.StatusUnprocessableEntity, "email is required")
	}

	if !govalidator.IsEmail(payload.Email) {
		API.ErrorResponse(ctx, http.StatusUnprocessableEntity, "invalid email format")
	}

	if !payload.IsPasswordSame() {
		API.ErrorResponse(ctx, http.StatusUnprocessableEntity, "confirmation password not same")
	}

	user, _ := userRepository.GetUserByUsername(payload.Username)
	if user != nil && govalidator.IsNotNull(user.Email) && govalidator.IsNotNull(user.Username) {
		API.ErrorResponse(ctx, http.StatusBadRequest, "user already registered")
	}

	uid, err := userRepository.CreateUser(&domain.User{
		Email:    payload.Email,
		Password: payload.Password,
		Phone:    payload.Phone,
		Username: payload.Username,
	})

	if err != nil {
		API.ErrorResponse(ctx, http.StatusInternalServerError, "Internal server error")
	}

	responsePayload := models.CreateUserResponse{
		ID:       uid,
		Username: payload.Username,
		Email:    payload.Email,
		Phone:    payload.Phone,
	}

	API.SuccessResponse(ctx, http.StatusCreated, "", responsePayload)
}
