package stdout

import (
	"github.com/IvanLutokhin/go-streamer/internal/pkg/log/logger/handler"
	"github.com/IvanLutokhin/go-streamer/internal/pkg/log/logger/handler/stream/writer"
	"io"
	"os"
)

func NewWriter(options handler.Options) (io.Writer, error) {
	return os.Stdout, nil
}

func ValidateWriter(options handler.Options) error {
	return nil
}

func init() {
	writer.RegisterWriter("stdout", NewWriter, ValidateWriter)
}
