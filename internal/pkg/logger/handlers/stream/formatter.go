package stream

import "github.com/IvanLutokhin/go-streamer/internal/pkg/logger"

type Formatter interface {
	Format(record logger.Record) ([]byte, error)
}
