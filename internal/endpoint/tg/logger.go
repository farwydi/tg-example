package tg

import "go.uber.org/zap"

func NewLogger(logger *zap.Logger) *Logger {
	return &Logger{
		SugaredLogger: logger.
			WithOptions(zap.AddCallerSkip(1)).
			Sugar(),
	}
}

type Logger struct {
	*zap.SugaredLogger
}

func (log Logger) Println(v ...interface{}) {
	log.Info(v)
}

func (log Logger) Printf(format string, v ...interface{}) {
	log.Infof(format, v)
}
