package API

import (
	"github.com/fiqrikm18/marketplace/core_services/pkg/models/API"
	"github.com/gin-gonic/gin"
)

func SuccessResponse(ctx *gin.Context, statusCode int, message string, data interface{}) {
	response := API.Response{
		Meta: API.ResponseMeta{
			Code:    statusCode,
			Status:  "ok",
			Message: message,
		},
		Data: data,
	}
	ctx.JSON(statusCode, response)
}

func ErrorResponse(ctx *gin.Context, statusCode int, message string) {
	response := API.Response{
		Meta: API.ResponseMeta{
			Code:    statusCode,
			Message: message,
			Status:  "failed",
		},
		Data: nil,
	}

	ctx.JSON(statusCode, response)
	ctx.Abort()
	return
}

func PaginateResponse() {

}
