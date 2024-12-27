package initialize

import (
	"go/go-backend-api/global"
	"go/go-backend-api/pkg/logger"

	"go.uber.org/zap"
)

func InitLogger() {
	global.Logger = logger.NewLogger()
	global.Logger.Info("init log success", zap.String("profile active", global.Profile))
}
