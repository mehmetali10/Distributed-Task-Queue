package loghelper

import (
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/log/level"
)

type Logger struct {
	logger log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		logger: log.NewJSONLogger(log.NewSyncWriter(os.Stdout)),
	}
}

func (l *Logger) Error(msg string, keyvals ...interface{}) {
	level.Error(l.logger).Log(keyvals...)
}

func (l *Logger) Info(msg string, keyvals ...interface{}) {
	level.Info(l.logger).Log(keyvals...)
}
