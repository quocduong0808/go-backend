package routers

import (
	"go/go-backend-api/internal/controller"

	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	v1.GET("/ping", controller.NewPongController().Pong)
	return r
}
