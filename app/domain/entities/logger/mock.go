package logHelper

type RepositoryMock struct {
	LogInfoFunc  func(operation string, msg string)
	LogWarnFunc  func(operation string, msg string)
	LogErrorFunc func(operation string, msg string)
}

func (r *RepositoryMock) LogInfo(operation string, msg string) {
}

func (r *RepositoryMock) LogWarn(operation string, msg string) {
}

func (r *RepositoryMock) LogError(operation string, msg string) {
}
