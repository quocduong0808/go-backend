//go:build wireinject
// +build wireinject

package wire

import (
	"go/go-backend-api/internal/controller"
	"go/go-backend-api/internal/repo"
	"go/go-backend-api/internal/service"

	"github.com/google/wire"
)

func InitUserHandler() (*controller.UserContoller, error) {
	wire.Build(controller.NewUserController, service.NewUserService, service.NewMailService, repo.NewUserRepository, repo.NewAuthRepository)
	return &controller.UserContoller{}, nil
}
