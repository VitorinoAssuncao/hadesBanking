package logHelper

import "context"

type LogHelper struct {
}

type Repository interface {
	LogInfo(ctx context.Context, msg string)
	LogWarn(ctx context.Context, msg string)
	LogError(ctx context.Context, msg string)
}
