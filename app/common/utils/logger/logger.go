package logHelper

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"stoneBanking/app/common/utils/config"
	logHelper "stoneBanking/app/domain/entities/logger"
	"stoneBanking/app/gateway/http/middleware"
)

type Log struct {
	logger    *zap.Logger
	requestID string
}

func NewLogger(config config.Config) logHelper.Logger {
	tempLogger := createLogger(config.Environment)
	logger := &Log{logger: tempLogger}
	return logger
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

func (l *Log) SetRequestIDFromContext(ctx context.Context) {
	requestID, _ := middleware.GetRequestIDFromContext(ctx)
	l.requestID = requestID
}

func (l Log) LogDebug(operation string, msg string) {
	l.logger.Debug(msg, zap.String("request_id", l.requestID), zap.String("operation", operation))
}

func (l Log) LogInfo(operation string, msg string) {
	l.logger.Info(msg, zap.String("request_id", l.requestID), zap.String("operation", operation))
}

func (l Log) LogWarn(operation string, msg string) {
	l.logger.Warn(msg, zap.String("request_id", l.requestID), zap.String("operation", operation))
}

func (l Log) LogError(operation string, msg string) {
	l.logger.Error(msg, zap.String("request_id", l.requestID), zap.String("operation", operation))
}
