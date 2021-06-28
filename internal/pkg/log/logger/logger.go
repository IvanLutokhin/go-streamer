package logger

import (
	"fmt"
	"github.com/IvanLutokhin/go-streamer/internal/pkg/log"
	"github.com/IvanLutokhin/go-streamer/internal/pkg/log/logger/handler"
	"os"
	"time"
)

type Logger struct {
	handlers []handler.Handler
}

func New(handlers ...handler.Handler) *Logger {
	return &Logger{handlers: handlers}
}

func (logger *Logger) Emergency(message string, fields ...log.Field) {
	logger.log(log.EMERGENCY, message, fields...)
}

func (logger *Logger) Alert(message string, fields ...log.Field) {
	logger.log(log.ALERT, message, fields...)
}

func (logger *Logger) Critical(message string, fields ...log.Field) {
	logger.log(log.CRITICAL, message, fields...)
}

func (logger *Logger) Error(message string, fields ...log.Field) {
	logger.log(log.ERROR, message, fields...)
}

func (logger *Logger) Warning(message string, fields ...log.Field) {
	logger.log(log.WARNING, message, fields...)
}

func (logger *Logger) Notice(message string, fields ...log.Field) {
	logger.log(log.NOTICE, message, fields...)
}

func (logger *Logger) Info(message string, fields ...log.Field) {
	logger.log(log.INFO, message, fields...)
}

func (logger *Logger) Debug(message string, fields ...log.Field) {
	logger.log(log.DEBUG, message, fields...)
}

func (logger *Logger) log(level log.Level, message string, fields ...log.Field) {
	if len(logger.handlers) == 0 {
		return
	}

	record := handler.Record{
		Timestamp: time.Now(),
		Caller:    handler.Caller(2),
		Level:     level,
		Message:   message,
		Fields:    fields,
	}

	for _, h := range logger.handlers {
		if h.IsHandling(record) {
			if err := h.Handle(record); err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "logger: %v\n", err)
			}
		}
	}
}
