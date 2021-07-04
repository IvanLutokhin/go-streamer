package logger

import (
	"github.com/IvanLutokhin/go-streamer/pkg/log"
	"time"
)

type Record struct {
	Timestamp time.Time
	Caller    Frame
	Level     Level
	Message   string
	Fields    []log.Field
}
