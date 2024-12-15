package initialize

import (
	"go/go-backend-api/global"
	"go/go-backend-api/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
