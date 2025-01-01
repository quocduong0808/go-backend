package user

import (
	"go/go-backend-api/internal/wire"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (ur *UserRouter) InitUserRouter(groupRouter *gin.RouterGroup) {
	publicUserRouter := groupRouter.Group("/user")
	pongController, _ := wire.InitPongHandler(0)
	userController, _ := wire.InitUserHandler()
	//user middleware
	//publicUserRouter.use(limit())
	{
		publicUserRouter.POST("/register", userController.Register)
		publicUserRouter.POST("/otp", pongController.Pong)
	}

	privateUserRouter := groupRouter.Group("/user")
	//user middleware
	//privateUserRouter.use(limit())
	//privateUserRouter.use(auth())
	//privateUserRouter.use(privileges())
	{
		privateUserRouter.GET("/info", pongController.Pong)
	}
}
