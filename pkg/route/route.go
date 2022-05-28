package route

import (
	"github.com/fiqrikm18/marketplace/core_services/pkg/route/API"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(srv *gin.Engine) {
	api := srv.Group("/api")
	API.APIRouter(api)
}
