package logHelper

import (
	"os"
	logHelper "stoneBanking/app/domain/entities/logger"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogRepository struct {
	logger     *zap.Logger
	enviroment string
}

func NewLogRepository() logHelper.Repository {
	logger := &LogRepository{enviroment: os.Getenv("ENVIROMENT")}
	logger.logger = createLogger(logger.enviroment)
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
	newLogger, _ := config.Build()
	return newLogger

}
func (l LogRepository) LogInfo(operation string, msg string) {
	l.logger.Info(msg, zap.String("operation:", operation), zap.String("enviroment:", l.enviroment))
}

func (l LogRepository) LogWarn(operation string, msg string) {
	l.logger.Warn(msg, zap.String("operation:", operation), zap.String("enviroment:", l.enviroment))
}

func (l LogRepository) LogError(operation string, msg string) {
	l.logger.Error(msg, zap.String("operation:", operation), zap.String("enviroment:", l.enviroment))
}
