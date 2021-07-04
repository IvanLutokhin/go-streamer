package logger

import (
	"github.com/IvanLutokhin/go-streamer/internal/app/streamer/config"
	"github.com/IvanLutokhin/go-streamer/internal/app/streamer/config/logger"
	"github.com/IvanLutokhin/go-streamer/pkg/log"

	_ "github.com/IvanLutokhin/go-streamer/internal/app/streamer/config/logger/handlers/stream"
	_ "github.com/IvanLutokhin/go-streamer/internal/app/streamer/config/logger/handlers/stream/formatters/text"
	_ "github.com/IvanLutokhin/go-streamer/internal/app/streamer/config/logger/handlers/stream/writers/stderr"
	_ "github.com/IvanLutokhin/go-streamer/internal/app/streamer/config/logger/handlers/stream/writers/stdout"
)

func New(config *config.Config) (log.Logger, error) {
	return logger.New(config.Logger)
}
