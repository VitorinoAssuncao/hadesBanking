package logHelper

import "context"

type RepositoryMock struct {
	LogInfoFunc                 func(operation string, msg string)
	LogWarnFunc                 func(operation string, msg string)
	LogErrorFunc                func(operation string, msg string)
	LogDebugFunc                func(operation string, msg string)
	SetRequestIDFromContextFunc func(ctx context.Context) Logger
}

func (r *RepositoryMock) LogInfo(operation string, msg string) {
}

func (r *RepositoryMock) LogDebug(operation string, msg string) {
}

func (r *RepositoryMock) LogWarn(operation string, msg string) {
}

func (r *RepositoryMock) LogError(operation string, msg string) {
}

func (r *RepositoryMock) SetRequestIDFromContext(ctx context.Context) Logger {
	return r
}
