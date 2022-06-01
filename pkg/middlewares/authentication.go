package middlewares

import (
	"github.com/asaskevich/govalidator"
	TokenStatus "github.com/fiqrikm18/marketplace/core_services/pkg/types/token/status"
	"github.com/fiqrikm18/marketplace/core_services/pkg/utils/API"
	"github.com/fiqrikm18/marketplace/core_services/pkg/utils/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.Request.Header["Authorization"]
		if len(authHeader) == 0 {
			API.ErrorResponse(context, http.StatusUnauthorized, "token not provided")
			return
		}

		if govalidator.IsNull(authHeader[0]) {
			API.ErrorResponse(context, http.StatusUnauthorized, "token not provided")
			return
		}

		tokenString := strings.Split(authHeader[0], " ")[1]
		if govalidator.IsEmail(tokenString) {
			API.ErrorResponse(context, http.StatusUnauthorized, "")
			return
		}

		tokenClaims, err := auth.ParseToken(tokenString)
		if err != nil {
			API.ErrorResponse(context, http.StatusInternalServerError, err.Error())
			return
		}

		tokenValidate := auth.CheckTokenExpired(time.Unix(tokenClaims.StandardClaims.ExpiresAt, 0))
		if tokenValidate == TokenStatus.TOKEN_STATUS_EXPIRED {
			API.ErrorResponse(context, http.StatusUnauthorized, "token expired")
			return
		}
	}
}
