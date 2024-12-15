package logger

import (
	"go/go-backend-api/pkg/setting"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewLogger(config setting.LoggerSetting) *zap.Logger {
	//fmt.Println("Log config info: ", config.Level, config.Filename, config.MaxSize, config.MaxBackups, config.MaxAge, config.Compress)
	logLevel := getLogLevel(config.Level)
	logFormat := getEncoderLog()
	logSync := zapcore.AddSync(&lumberjack.Logger{
		Filename:   config.Filename,
		MaxSize:    config.MaxSize, // megabytes
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,   //days
		Compress:   config.Compress, // disabled by default
	})
	logConsleSync := zapcore.AddSync(os.Stderr)
	core := zapcore.NewCore(logFormat, zapcore.NewMultiWriteSyncer(logSync, logConsleSync), logLevel)
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
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.TimeKey = "time"
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encodeConfig)
}
