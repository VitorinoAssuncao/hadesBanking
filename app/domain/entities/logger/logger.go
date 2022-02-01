package logHelper

import "context"

type Logger interface {
	LogInfo(operation string, msg string)
	LogWarn(operation string, msg string)
	LogError(operation string, msg string)
	SetRequestIDFromContext(ctx context.Context) Logger
}
