package logger

import (
	"github.com/IvanLutokhin/go-streamer/internal/app/streamer/config"
	"github.com/IvanLutokhin/go-streamer/internal/pkg/log"
	"github.com/IvanLutokhin/go-streamer/internal/pkg/log/logger"
	"github.com/IvanLutokhin/go-streamer/internal/pkg/log/logger/handler"

	_ "github.com/IvanLutokhin/go-streamer/internal/pkg/log/logger/handler/stream"
	_ "github.com/IvanLutokhin/go-streamer/internal/pkg/log/logger/handler/stream/formatter/text"
	_ "github.com/IvanLutokhin/go-streamer/internal/pkg/log/logger/handler/stream/writer/stderr"
	_ "github.com/IvanLutokhin/go-streamer/internal/pkg/log/logger/handler/stream/writer/stdout"
)

func New(config *config.Config) (log.Logger, error) {
	var handlers []handler.Handler
	for _, handlerConfig := range config.Logger.Handlers {
		if !handlerConfig.Enabled {
			continue
		}

		h, err := handler.NewHandler(handlerConfig.Code, handlerConfig.Options)
		if err != nil {
			return nil, err
		}

		handlers = append(handlers, h)
	}

	return logger.New(handlers...), nil
}
