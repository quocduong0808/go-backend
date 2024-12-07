package routers

import (
	"go/go-backend-api/internal/controller"
	"go/go-backend-api/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.AuthMiddleware())
	v1 := r.Group("/api/v1")
	v1.GET("/ping", controller.NewPongController().Pong)
	return r
}
