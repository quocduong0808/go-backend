package user

import (
	"go/go-backend-api/internal/controller"

	"github.com/gin-gonic/gin"
)

type ProductRouter struct {
}

func (pr *ProductRouter) InitProductRouter(groupRouter *gin.RouterGroup) {
	publicProductRouter := groupRouter.Group("/product")
	//user middleware
	//publicProductRouter.use(limit())
	{
		publicProductRouter.POST("/search", controller.NewPongController().Pong)
		publicProductRouter.POST("/detail/:id", controller.NewPongController().Pong)
	}
}
