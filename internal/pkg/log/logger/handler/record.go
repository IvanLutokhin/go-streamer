package handler

import (
	"github.com/IvanLutokhin/go-streamer/internal/pkg/log"
	"time"
)

type Record struct {
	Timestamp time.Time
	Caller    Frame
	Level     log.Level
	Message   string
	Fields    []log.Field
}
