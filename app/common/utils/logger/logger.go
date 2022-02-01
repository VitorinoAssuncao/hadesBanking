package logHelper

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"stoneBanking/app/common/utils/config"
	logHelper "stoneBanking/app/domain/entities/logger"
)

type Log struct {
	logger *zap.Logger
}

func NewLogger(config config.Config) logHelper.Logger {
	tempLogger := createLogger(config.Environment)
	logger := &Log{tempLogger}
	return logger
}

func SetTracerID(l Log, requestID string) *Log {
	newLogger := l.logger.With(zap.String("requestID", requestID))
	newLog := &Log{logger: newLogger}
	return newLog
}

func createLogger(env string) *zap.Logger {
	config := zap.NewProductionConfig()
	if env == "development" {
		config = zap.NewDevelopmentConfig()
	}
	config.DisableCaller = true
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.DisableStacktrace = true
	logger, _ := config.Build()
	newLogger := logger.With(zap.String("environment", env))
	return newLogger

}
func (l Log) LogInfo(operation string, msg string) {
	l.logger.Info(msg, zap.String("operation:", operation))
}

func (l Log) LogWarn(operation string, msg string) {
	l.logger.Warn(msg, zap.String("operation:", operation))
}

func (l Log) LogError(operation string, msg string) {
	l.logger.Error(msg, zap.String("operation:", operation))
}
