package routers

import (
	"go/go-backend-api/internal/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	v1.GET("/ping", controller.NewPongController().Pong)
	return r
}

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong..pong",
	})
}
