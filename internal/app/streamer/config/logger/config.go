package logger

type Config struct {
	Handlers []HandlerConfig `json:"handlers"`
}

func NewDefaultConfig() *Config {
	return &Config{
		Handlers: []HandlerConfig{
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
