package stream

import (
	"encoding/json"
	loggerConfig "github.com/IvanLutokhin/go-streamer/internal/app/streamer/config/logger"
	"github.com/IvanLutokhin/go-streamer/internal/pkg/logger"
	"github.com/IvanLutokhin/go-streamer/internal/pkg/logger/handlers/stream"
	"io"
)

func NewHandler(options map[string]interface{}) (h logger.Handler, err error) {
	var handlerConfig *HandlerConfig
	if handlerConfig, err = ParseHandlerConfig(options); err != nil {
		return
	}

	var level logger.Level
	if level, err = logger.LevelFromString(handlerConfig.Level); err != nil {
		return
	}

	var f stream.Formatter
	if f, err = NewFormatter(handlerConfig.Formatter.Code, handlerConfig.Formatter.Options); err != nil {
		return
	}

	var w io.Writer
	if w, err = NewWriter(handlerConfig.Writer.Code, handlerConfig.Writer.Options); err != nil {
		return
	}

	h = &stream.Handler{
		Level:     level,
		Formatter: f,
		Writer:    w,
	}

	return
}

func ValidateHandler(options map[string]interface{}) (err error) {
	var handlerConfig *HandlerConfig
	if handlerConfig, err = ParseHandlerConfig(options); err != nil {
		return
	}

	if _, err = logger.LevelFromString(handlerConfig.Level); err != nil {
		return
	}

	if err = ValidateFormatter(handlerConfig.Formatter.Code, handlerConfig.Formatter.Options); err != nil {
		return
	}

	if err = ValidateWriter(handlerConfig.Writer.Code, handlerConfig.Writer.Options); err != nil {
		return
	}

	return
}

func ParseHandlerConfig(options map[string]interface{}) (*HandlerConfig, error) {
	handlerConfig := NewDefaultHandlerConfig()

	if options == nil {
		return handlerConfig, nil
	}

	data, err := json.Marshal(options)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, handlerConfig); err != nil {
		return nil, err
	}

	return handlerConfig, nil
}

func init() {
	loggerConfig.RegisterHandler("stream", NewHandler, ValidateHandler)
}
