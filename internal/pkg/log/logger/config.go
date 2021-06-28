package logger

import "github.com/IvanLutokhin/go-streamer/internal/pkg/log/logger/handler"

type Config struct {
	Handlers []handler.Config `json:"handlers"`
}

func NewDefaultConfig() *Config {
	return &Config{
		Handlers: []handler.Config{
			{
				Code:    "stream",
				Enabled: true,
				Options: nil, // use default
			},
		},
	}
}

func (config *Config) Validate() (err error) {
	for _, handlerConfig := range config.Handlers {
		if err = handlerConfig.Validate(); err != nil {
			return
		}
	}

	return
}
