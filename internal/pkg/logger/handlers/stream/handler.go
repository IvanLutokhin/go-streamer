package stream

import (
	"github.com/IvanLutokhin/go-streamer/internal/pkg/logger"
	"io"
)

type Handler struct {
	Level     logger.Level
	Formatter Formatter
	Writer    io.Writer
}

func (handler *Handler) IsHandling(record logger.Record) bool {
	return handler.Level >= record.Level
}

func (handler *Handler) Handle(record logger.Record) error {
	bytes, err := handler.Formatter.Format(record)
	if err != nil {
		return err
	}

	if _, err = handler.Writer.Write(bytes); err != nil {
		return err
	}

	return nil
}
