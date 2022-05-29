package auth

import (
	"github.com/asaskevich/govalidator"
	"github.com/fiqrikm18/marketplace/core_services/pkg/config"
	"github.com/fiqrikm18/marketplace/core_services/pkg/domain"
	"github.com/fiqrikm18/marketplace/core_services/pkg/models"
	"github.com/fiqrikm18/marketplace/core_services/pkg/repositories"
	"github.com/fiqrikm18/marketplace/core_services/pkg/utils/API"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

var (
	userRepository *repositories.UserRepository
)

func init() {
	dbConf, err := config.NewConnection()
	if err != nil {
		panic(err)
	}

	userRepository = repositories.NewUserRepository(dbConf.DB)
}

func Login(ctx *gin.Context) {

}

func Logout(ctx *gin.Context) {

}

func Register(ctx *gin.Context) {
	var payload models.CreateUserRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		API.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	phoneRegex, err := regexp.Compile("\\d+")
	if err != nil {
		API.ErrorResponse(ctx, http.StatusInternalServerError, "internal server error")
		return
	}

	if !phoneRegex.MatchString(payload.Phone) {
		API.ErrorResponse(ctx, http.StatusUnprocessableEntity, "phone must be numeric")
		return
	}

	if govalidator.IsNull(payload.Email) {
		API.ErrorResponse(ctx, http.StatusUnprocessableEntity, "email is required")
		return
	}

	if !govalidator.IsEmail(payload.Email) {
		API.ErrorResponse(ctx, http.StatusUnprocessableEntity, "invalid email format")
		return
	}

	if !payload.IsPasswordSame() {
		API.ErrorResponse(ctx, http.StatusUnprocessableEntity, "confirmation password not same")
		return
	}

	user, err := userRepository.ValidateUserExist(payload.Username, payload.Email)
	if err != nil {
		API.ErrorResponse(ctx, http.StatusInternalServerError, "internal server error")
		return
	}

	if govalidator.IsNotNull(user.Email) && govalidator.IsNotNull(user.Username) {
		API.ErrorResponse(ctx, http.StatusBadRequest, "user already registered")
		return
	}

	uid, err := userRepository.CreateUser(&domain.User{
		Email:    payload.Email,
		Password: payload.Password,
		Phone:    payload.Phone,
		Username: payload.Username,
	})

	if err != nil {
		API.ErrorResponse(ctx, http.StatusInternalServerError, "Internal server error")
		return
	}

	responsePayload := models.CreateUserResponse{
		ID:       uid,
		Username: payload.Username,
		Email:    payload.Email,
		Phone:    payload.Phone,
	}

	API.SuccessResponse(ctx, http.StatusCreated, "", responsePayload)
}
