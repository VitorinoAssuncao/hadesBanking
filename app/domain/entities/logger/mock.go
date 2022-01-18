package logHelper

import "context"

type RepositoryMock struct {
	LogInfoFunc  func(ctx context.Context, msg string)
	LogWarnFunc  func(ctx context.Context, msg string)
	LogErrorFunc func(ctx context.Context, msg string)
}

func (r *RepositoryMock) LogInfo(ctx context.Context, msg string) {
}

func (r *RepositoryMock) LogWarn(ctx context.Context, msg string) {
}

func (r *RepositoryMock) LogError(ctx context.Context, msg string) {
}
