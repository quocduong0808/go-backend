package global

import (
	"go/go-backend-api/pkg/setting"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *zap.Logger
	MyDB   *gorm.DB
)

func HandleErrorPanic(err error, msg string) {
	if err != nil {
		Logger.Error(msg, zap.Error(err))
		panic(msg)
	}
}
