package logger

import (
	"go/go-backend-api/global"
	"go/go-backend-api/global/consts"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewLogger() *zap.Logger {
	//fmt.Println("Log logConfig info: ", logConfig.Level, logConfig.Filename, logConfig.MaxSize, logConfig.MaxBackups, logConfig.MaxAge, logConfig.Compress)
	logConfig := global.Config.Logger
	profile := global.Config.Server.Profile
	logLevel := getLogLevel(logConfig.Level)
	logFormat := getEncoderLog()
	logSync := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logConfig.Filename,
		MaxSize:    logConfig.MaxSize, // megabytes
		MaxBackups: logConfig.MaxBackups,
		MaxAge:     logConfig.MaxAge,   //days
		Compress:   logConfig.Compress, // disabled by default
	})
	var core zapcore.Core
	switch profile {
	case consts.PROFILE_DEV:
		logConsleSync := zapcore.AddSync(os.Stderr)
		core = zapcore.NewCore(logFormat, zapcore.NewMultiWriteSyncer(logSync, logConsleSync), logLevel)
	case consts.PROFILE_PROD:
		core = zapcore.NewCore(logFormat, logSync, logLevel)
	default:
		panic("error profile when init logger")
	}

	logger := zap.New(core, zap.AddCaller())
	return logger
	// logger.Info("Info log", zap.Int("line", 1))
	// logger.Error("Info error", zap.Int("line", 2))
}

func getLogLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

// format log
func getEncoderLog() zapcore.Encoder {
	encodelogConfig := zap.NewProductionEncoderConfig()
	encodelogConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodelogConfig.TimeKey = "time"
	encodelogConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encodelogConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encodelogConfig)
}
