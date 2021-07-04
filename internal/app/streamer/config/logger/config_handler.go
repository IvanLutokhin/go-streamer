package logger

type HandlerConfig struct {
	Enabled bool                   `json:"enabled"`
	Code    string                 `json:"code"`
	Options map[string]interface{} `json:"options"`
}

func (config *HandlerConfig) Validate() error {
	if !config.Enabled {
		return nil
	}

	return ValidateHandler(config.Code, config.Options)
}
