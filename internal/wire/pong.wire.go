//go:build wireinject
// +build wireinject

package wire

import (
	"go/go-backend-api/internal/controller"
	"go/go-backend-api/internal/repo"
	"go/go-backend-api/internal/service"

	"github.com/google/wire"
)

func InitPongHandler(id int) (*controller.PongController, error) {
	wire.Build(controller.NewPongController, service.NewPongService, repo.NewPongRepo)
	return &controller.PongController{}, nil
}
