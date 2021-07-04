package logger

import "github.com/IvanLutokhin/go-streamer/internal/pkg/logger"

func New(config *Config) (*logger.Logger, error) {
	var handlers []logger.Handler
	for _, handlerConfig := range config.Handlers {
		if !handlerConfig.Enabled {
			continue
		}

		handler, err := NewHandler(handlerConfig.Code, handlerConfig.Options)
		if err != nil {
			return nil, err
		}

		handlers = append(handlers, handler)
	}

	return logger.New(handlers...), nil
}
