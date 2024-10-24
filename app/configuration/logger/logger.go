package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Logger *zap.Logger
	err    error
)

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{"/app/logs/logs.log", "stdout"},
		Level:       zap.NewAtomicLevel(),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			LevelKey:     "level",
			TimeKey:      "datetime",
			FunctionKey:  "function",
			CallerKey:    "caller",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.FullCallerEncoder,
		},
	}

	Logger, err = logConfig.Build()
	if err != nil {
		panic(err) // Captura e reporta se houve erro na criação do logger
	}
}
