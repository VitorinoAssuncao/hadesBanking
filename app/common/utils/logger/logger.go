package logHelper

import (
	logHelper "stoneBanking/app/domain/entities/logger"

	"go.uber.org/zap"
)

type LogRepository struct {
	logger *zap.Logger
}

func NewLogRepository() logHelper.Repository {
	newLogger, _ := zap.NewDevelopment(zap.AddCaller())
	return &LogRepository{
		logger: newLogger,
	}
}

func (l LogRepository) LogInfo(operation string, msg string) {
	l.logger.Info(operation, zap.String("message:", msg))
	l.logger.Sync()
}

func (l LogRepository) LogWarn(operation string, msg string) {
	l.logger.Warn(operation, zap.String("message:", msg))
	l.logger.Sync()
}

func (l LogRepository) LogError(operation string, msg string) {
	l.logger.Error(operation, zap.String("message:", msg))
	l.logger.Sync()
}
