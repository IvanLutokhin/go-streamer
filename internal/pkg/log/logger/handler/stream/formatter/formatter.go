package formatter

import "github.com/IvanLutokhin/go-streamer/internal/pkg/log/logger/handler"

type Formatter interface {
	Format(record handler.Record) ([]byte, error)
}
