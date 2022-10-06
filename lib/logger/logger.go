package logger

import (
	"sync"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger
}

type Config struct {
	Level   string `mapstructure:"level"`
	Service string `mapstructure:"service"`
}

var initOnce sync.Once
var logger, _ = zap.NewDevelopment()

func InitLogger(cf *Config) {
	initOnce.Do(func() {
		if cf.Service == "" {
			logger.Fatal("Service name is not set")
		}

		cfg := zap.NewDevelopmentConfig()
		switch cf.Level {
		case "debug":
			cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		case "info":
			cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
		case "warn":
			cfg.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
		case "error":
			cfg.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
		case "fatal":
			cfg.Level = zap.NewAtomicLevelAt(zap.FatalLevel)
		default:
			cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
		}
		cfg.OutputPaths = []string{"stdout"}

		l, err := cfg.Build()
		if err != nil {
			logger.Fatal(errors.WithStack(err).Error())
		}

		logger = l.With(zap.String("service", cf.Service))
		logger.Info("Logger initialized")
	})
}

func Info(msg string, fields ...zapcore.Field) {
	logger.WithOptions(zap.AddCallerSkip(1)).Info(msg, fields...)
}

func Warn(msg string, fields ...zapcore.Field) {
	logger.WithOptions(zap.AddCallerSkip(1)).Warn(msg, fields...)
}

func Debug(msg string, fields ...zapcore.Field) {
	logger.WithOptions(zap.AddCallerSkip(1)).Debug(msg, fields...)
}

func Fatal(msg string, fields ...zapcore.Field) {
	logger.WithOptions(zap.AddCallerSkip(1)).Fatal(msg, fields...)
}

func Panic(msg string, fields ...zapcore.Field) {
	logger.WithOptions(zap.AddCallerSkip(1)).Panic(msg, fields...)
}

func Error(msg string, fields ...zapcore.Field) {
	logger.WithOptions(zap.AddCallerSkip(1)).Error(msg, fields...)
}

func With(fields ...zapcore.Field) *zap.Logger {
	return logger.With(fields...)
}
