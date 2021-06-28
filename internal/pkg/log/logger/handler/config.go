package handler

type Config struct {
	Enabled bool    `json:"enabled"`
	Code    string  `json:"code"`
	Options Options `json:"options"`
}

func (config *Config) Validate() error {
	if !config.Enabled {
		return nil
	}

	return ValidateHandler(config.Code, config.Options)
}
