package logHelper

type LogHelper struct {
}

type Logger interface {
	LogInfo(operation string, msg string)
	LogWarn(operation string, msg string)
	LogError(operation string, msg string)
}
