package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router(v1 *gin.RouterGroup) {
	v1.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Hello, This is the marketplace API",
		})
	})
}
