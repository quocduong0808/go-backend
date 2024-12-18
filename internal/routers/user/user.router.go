package user

import (
	"go/go-backend-api/internal/controller"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (ur *UserRouter) InitUserRouter(groupRouter *gin.RouterGroup) {
	publicUserRouter := groupRouter.Group("/user")
	//user middleware
	//publicUserRouter.use(limit())
	{
		publicUserRouter.POST("/register", controller.NewPongController().Pong)
		publicUserRouter.POST("/otp", controller.NewPongController().Pong)
	}

	privateUserRouter := groupRouter.Group("/user")
	//user middleware
	//privateUserRouter.use(limit())
	//privateUserRouter.use(auth())
	//privateUserRouter.use(privileges())
	{
		privateUserRouter.GET("/info", controller.NewPongController().Pong)
	}
}
