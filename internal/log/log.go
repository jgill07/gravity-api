package log

import (
	"go.uber.org/zap"
)

var log *zap.Logger

func InitLogger() {
	var cfg zap.Config = zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		Encoding:         "json",
		Development:      false,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	log = zap.Must(cfg.Build())
	defer log.Sync()
}

func l() *zap.Logger {
	if log == nil {
		InitLogger()
	}
	return log
}

func with(fields ...zap.Field) *zap.Logger {
	return l().With(fields...)
}

func WithError(err error) *zap.Logger {
	return with(zap.Error(err))
}

func WithFields(fields ...zap.Field) *zap.Logger {
	return with(fields...)
}

func Debug(msg string) {
	l().Debug(msg)
}

func Error(msg string) {
	l().Error(msg)
}

func Info(msg string) {
	l().Info(msg)
}
