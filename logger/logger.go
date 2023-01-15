package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func CreateLogger(path string) *zap.Logger {
	consoleErrors := zapcore.Lock(os.Stdout)
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	core := zapcore.NewCore(consoleEncoder, consoleErrors, zapcore.InfoLevel)
	return zap.New(core, zap.Fields(zap.String("path", path)))
}
