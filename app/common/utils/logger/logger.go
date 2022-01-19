package logHelper

import (
	"os"
	logHelper "stoneBanking/app/domain/entities/logger"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogRepository struct {
	logger *zap.Logger
}

func NewLogRepository() logHelper.Repository {
	enviroment := os.Getenv("ENVIROMENT")
	tempLogger := createLogger(enviroment)
	logger := &LogRepository{tempLogger}
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
	newLogger := logger.With(zap.String("enviroment", env))
	return newLogger

}
func (l LogRepository) LogInfo(operation string, msg string) {
	l.logger.Info(msg, zap.String("operation:", operation))
	l.logger.With()
}

func (l LogRepository) LogWarn(operation string, msg string) {
	l.logger.Warn(msg, zap.String("operation:", operation))
}

func (l LogRepository) LogError(operation string, msg string) {
	l.logger.Error(msg, zap.String("operation:", operation))
}
