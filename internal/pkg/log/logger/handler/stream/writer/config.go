package writer

import "github.com/IvanLutokhin/go-streamer/internal/pkg/log/logger/handler"

type Config struct {
	Code    string          `json:"code"`
	Options handler.Options `json:"options"`
}
