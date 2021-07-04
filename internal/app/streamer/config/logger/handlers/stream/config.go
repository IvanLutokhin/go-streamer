package stream

type HandlerConfig struct {
	Level     string           `json:"level"`
	Formatter *FormatterConfig `json:"formatter"`
	Writer    *WriterConfig    `json:"writer"`
}

func NewDefaultHandlerConfig() *HandlerConfig {
	return &HandlerConfig{
		Level: "info",
		Formatter: &FormatterConfig{
			Code:    "text",
			Options: nil, // use default
		},
		Writer: &WriterConfig{
			Code:    "stdout",
			Options: nil, // use default
		},
	}
}
