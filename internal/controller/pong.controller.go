package controller

import (
	"go/go-backend-api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct {
	pongService *service.PongService
}

func NewPongController() *PongController {
	return &PongController{
		pongService: service.NewPongService(),
	}
}

func (pc *PongController) Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": pc.pongService.Pong(),
	})
}
