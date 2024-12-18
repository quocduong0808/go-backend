package manage

import (
	"go/go-backend-api/internal/controller"

	"github.com/gin-gonic/gin"
)

type AdminRouter struct {
}

func (ar *AdminRouter) InitAdminRouter(groupRouter *gin.RouterGroup) {
	publicAdminRouter := groupRouter.Group("/admin")
	//user middleware
	//publicAdminRouter.use(limit())
	{
		publicAdminRouter.POST("/login", controller.NewPongController().Pong)
	}

	privateAdminRouter := groupRouter.Group("/admin")
	//user middleware
	//privateAdminRouter.use(limit())
	//privateAdminRouter.use(auth())
	//privateAdminRouter.use(privileges())
	{
		privateAdminRouter.GET("/active_user", controller.NewPongController().Pong)
	}
}
