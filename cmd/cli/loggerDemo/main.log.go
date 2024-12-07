package main

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	logFormat := getEncoderLog()
	logSync := getLogWriterSync()
	core := zapcore.NewCore(logFormat, logSync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())
	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Info error", zap.Int("line", 2))
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

// create log file
func getLogWriterSync() zapcore.WriteSyncer {
	file, err := os.OpenFile("./logs/app.log", os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(fmt.Errorf("error when create log writer %w", err))
	}
	synFile := zapcore.AddSync(file)
	synConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(synFile, synConsole)
}
