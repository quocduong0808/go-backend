package manage

import (
	"go/go-backend-api/internal/wire"

	"github.com/gin-gonic/gin"
)

type AdminRouter struct {
}

func (ar *AdminRouter) InitAdminRouter(groupRouter *gin.RouterGroup) {
	publicAdminRouter := groupRouter.Group("/admin")
	pongController, _ := wire.InitPongHandler(0)
	//user middleware
	//publicAdminRouter.use(limit())
	{
		publicAdminRouter.POST("/login", pongController.Pong)
	}

	privateAdminRouter := groupRouter.Group("/admin")
	//user middleware
	//privateAdminRouter.use(limit())
	//privateAdminRouter.use(auth())
	//privateAdminRouter.use(privileges())
	{
		privateAdminRouter.GET("/active_user", pongController.Pong)
	}
}
