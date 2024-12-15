package global

import (
	"go/go-backend-api/pkg/setting"

	"go.uber.org/zap"
)

var (
	Config setting.Config
	Logger *zap.Logger
)
