package adapters

import (
	"main/internal/ports"

	"go.uber.org/zap"
)

type ZapLogger struct {
	logger *zap.Logger
}

func (l *ZapLogger) NewZapLogger(logger *zap.Logger) ports.ZapLogger {
	return &ZapLogger{
		logger: logger,
	}
}

func (l *ZapLogger) Info(msg string) {
	l.logger.Info(msg)
}

func (l *ZapLogger) Error(msg string) {
	l.logger.Error(msg)
}

func (l *ZapLogger) Debug(msg string) {
	l.logger.Debug(msg)
}
