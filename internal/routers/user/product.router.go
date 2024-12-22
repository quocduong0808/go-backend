package user

import (
	"go/go-backend-api/internal/wire"

	"github.com/gin-gonic/gin"
)

type ProductRouter struct {
}

func (pr *ProductRouter) InitProductRouter(groupRouter *gin.RouterGroup) {
	publicProductRouter := groupRouter.Group("/product")
	pongController, _ := wire.InitPongHandler(0)
	//user middleware
	//publicProductRouter.use(limit())
	{
		publicProductRouter.POST("/search", pongController.Pong)
		publicProductRouter.POST("/detail/:id", pongController.Pong)
	}
}
