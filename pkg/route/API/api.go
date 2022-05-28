package API

import (
	v12 "github.com/fiqrikm18/marketplace/core_services/pkg/route/API/v1"
	"github.com/gin-gonic/gin"
)

func APIRouter(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	v12.Router(v1)
}
