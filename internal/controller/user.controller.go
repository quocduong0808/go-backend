package controller

import (
	"go/go-backend-api/internal/service"
	"go/go-backend-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserContoller struct {
	userService service.IUserService
}

func (uc *UserContoller) Register(c *gin.Context) {
	messageCode := uc.userService.Register("quocduong0897@gmail.com", "123456", "name", "password")
	response.ResponseSuccess(c, messageCode, response.MSG[messageCode], nil)
}

func NewUserController(userService service.IUserService) *UserContoller {
	return &UserContoller{
		userService: userService,
	}
}
