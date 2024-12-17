package global

import (
	"go/go-backend-api/pkg/setting"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *zap.Logger
	MyDB   *gorm.DB
	Redis  *redis.Client
)

func HandleErrorPanic(err error, msg string) {
	if err != nil {
		Logger.Error(msg, zap.Error(err))
		panic(msg)
	}
}
