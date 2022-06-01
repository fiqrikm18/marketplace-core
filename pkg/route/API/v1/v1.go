package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"

	AuthController "github.com/fiqrikm18/marketplace/core_services/pkg/controllers/auth"
)

func Router(v1 *gin.RouterGroup) {
	v1.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Hello, This is the marketplace API",
		})
	})

	// authentication
	auth := v1.Group("/auth")
	{
		auth.POST("/register", AuthController.Register)
		auth.POST("/login", AuthController.Login)

		// TODO: add middleware
		auth.GET("/logout", AuthController.Logout)
	}
}
