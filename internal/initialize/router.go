package initialize

import (
	"go/go-backend-api/internal/routers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := routers.CreateRouter()
	return r
}
