package stdout

import (
	handlerConfig "github.com/IvanLutokhin/go-streamer/internal/app/streamer/config/logger/handlers/stream"
	"io"
	"os"
)

func NewWriter(map[string]interface{}) (io.Writer, error) {
	return os.Stderr, nil
}

func ValidateWriter(map[string]interface{}) error {
	return nil
}

func init() {
	handlerConfig.RegisterWriter("stderr", NewWriter, ValidateWriter)
}
