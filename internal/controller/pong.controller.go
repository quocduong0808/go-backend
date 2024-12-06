package controller

import (
	"go/go-backend-api/internal/service"
	"go/go-backend-api/pkg/response"

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
	response.ResponseSuccess(c, response.SuccessCode, response.MSG[response.SuccessCode], pc.pongService.Pong())
}
