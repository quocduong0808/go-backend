package routers

import (
	"go/go-backend-api/global"
	"go/go-backend-api/global/consts"
	"go/go-backend-api/internal/wire"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	var r *gin.Engine
	switch global.Profile {
	case consts.PROFILE_DEV, consts.PROFILE_LOCAL:
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	case consts.PROFILE_PROD:
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	default:
		panic("cannot get profile to init router")
	}
	//r.Use(logger)
	//r.Use(cross)
	//r.Use(limit)
	rootGroup := r.Group("/api/v1")
	{
		pongController, _ := wire.InitPongHandler(0)
		rootGroup.GET("/get-status", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, map[string]string{"key": "value"})
		})
		rootGroup.GET("/ping", pongController.Pong)
	}
	{
		AppRouter.Admin.InitAdminRouter(rootGroup)
	}
	{
		AppRouter.User.InitUserRouter(rootGroup)
		AppRouter.User.InitProductRouter(rootGroup)
	}

	//v1.GET("/ping", controller.NewPongController().Pong)
	return r
}
