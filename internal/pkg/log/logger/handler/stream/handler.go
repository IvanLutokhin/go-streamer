package stream

import (
	"github.com/IvanLutokhin/go-streamer/internal/pkg/log"
	"github.com/IvanLutokhin/go-streamer/internal/pkg/log/logger/handler"
	"github.com/IvanLutokhin/go-streamer/internal/pkg/log/logger/handler/stream/formatter"
	"github.com/IvanLutokhin/go-streamer/internal/pkg/log/logger/handler/stream/writer"
	"io"
)

type Handler struct {
	level     log.Level
	formatter formatter.Formatter
	writer    io.Writer
}

func NewHandler(options handler.Options) (h handler.Handler, err error) {
	var handlerOptions *HandlerOptions
	if handlerOptions, err = ParseHandlerOptions(options); err != nil {
		return
	}

	var level log.Level
	if level, err = log.LevelFromString(handlerOptions.Level); err != nil {
		return
	}

	var f formatter.Formatter
	if f, err = formatter.NewFormatter(handlerOptions.Formatter.Code, handlerOptions.Formatter.Options); err != nil {
		return
	}

	var w io.Writer
	if w, err = writer.NewWriter(handlerOptions.Writer.Code, handlerOptions.Writer.Options); err != nil {
		return
	}

	h = &Handler{
		level:     level,
		formatter: f,
		writer:    w,
	}

	return
}

func ValidateHandler(options handler.Options) (err error) {
	var handlerOptions *HandlerOptions
	if handlerOptions, err = ParseHandlerOptions(options); err != nil {
		return
	}

	if _, err = log.LevelFromString(handlerOptions.Level); err != nil {
		return
	}

	if err = formatter.ValidateFormatter(handlerOptions.Formatter.Code, handlerOptions.Formatter.Options); err != nil {
		return
	}

	if err = writer.ValidateWriter(handlerOptions.Writer.Code, handlerOptions.Writer.Options); err != nil {
		return
	}

	return
}

func ParseHandlerOptions(options handler.Options) (*HandlerOptions, error) {
	handlerOptions := NewDefaultHandlerOptions()

	if err := options.Unmarshal(handlerOptions); err != nil {
		return nil, err
	}

	return handlerOptions, nil
}

func (handler *Handler) IsHandling(record handler.Record) bool {
	return handler.level >= record.Level
}

func (handler *Handler) Handle(record handler.Record) error {
	bytes, err := handler.formatter.Format(record)
	if err != nil {
		return err
	}

	if _, err = handler.writer.Write(bytes); err != nil {
		return err
	}

	return nil
}

func init() {
	handler.RegisterHandler("stream", NewHandler, ValidateHandler)
}
