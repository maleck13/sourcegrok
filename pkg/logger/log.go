package logger

import (
	"github.com/Sirupsen/logrus"
	"github.com/maleck13/sourcegrok/pkg/sourcegrok"
)

type Logger struct {
	logger *logrus.Logger
}

func (l *Logger) Log(level int, messages ...interface{}) {
	if level == sourcegrok.LogLevelDebug {
		l.logger.Debug(messages)
	}
	if level == sourcegrok.LogLevelInfo {
		l.logger.Info(messages)
	}
	if level == sourcegrok.LogLevelError {
		l.logger.Error(messages)
	}
}

// New returns a new Logger implements Logger interface
func New(l *logrus.Logger) *Logger {
	return &Logger{
		logger: l,
	}
}
