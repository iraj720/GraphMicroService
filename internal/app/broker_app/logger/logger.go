package logger

import (
	"github.com/sirupsen/logrus"
)

type GraphLogger interface {
	Log(args ...interface{})
}

type graphLogger struct {
	logger *logrus.Logger
}

func NewGraphLogger(logger *logrus.Logger) GraphLogger {
	return &graphLogger{logger: logger}
}

func (gl *graphLogger) Log(args ...interface{}) {
	gl.logger.Info(args...)
}
