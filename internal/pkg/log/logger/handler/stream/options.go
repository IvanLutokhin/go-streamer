package stream

import (
	"github.com/IvanLutokhin/go-streamer/internal/pkg/log/logger/handler/stream/formatter"
	"github.com/IvanLutokhin/go-streamer/internal/pkg/log/logger/handler/stream/writer"
)

type HandlerOptions struct {
	Level     string            `json:"level"`
	Formatter *formatter.Config `json:"formatter"`
	Writer    *writer.Config    `json:"writer"`
}

func NewDefaultHandlerOptions() *HandlerOptions {
	return &HandlerOptions{
		Level: "info",
		Formatter: &formatter.Config{
			Code:    "text",
			Options: nil, // use default
		},
		Writer: &writer.Config{
			Code:    "stdout",
			Options: nil, // use default
		},
	}
}
