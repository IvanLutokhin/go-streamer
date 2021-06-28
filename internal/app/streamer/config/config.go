package config

import (
	"encoding/json"
	"fmt"
	"github.com/IvanLutokhin/go-streamer/internal/pkg/log/logger"
	"os"
)

type Config struct {
	Logger *logger.Config `json:"logger"`
}

func NewDefault() *Config {
	return &Config{
		Logger: logger.NewDefaultConfig(),
	}
}

func Load(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("config: %v", err)
	}

	defer file.Close()

	config := NewDefault()

	if err := json.NewDecoder(file).Decode(config); err != nil {
		return nil, fmt.Errorf("config: %v", err)
	}

	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("config: %v", err)
	}

	return config, nil
}

func (config *Config) Validate() (err error) {
	if err = config.Logger.Validate(); err != nil {
		return
	}

	return
}
