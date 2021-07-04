package text

import (
	"bytes"
	"errors"
	"github.com/IvanLutokhin/go-streamer/internal/pkg/logger"
	"github.com/IvanLutokhin/go-streamer/pkg/log"
	"sync"
	"testing"
	"time"
)

func TestFormatter_Format(t *testing.T) {
	template, err := ParseTemplate("%datetime% %caller% %level% %message% %context%", false)
	if err != nil {
		t.Error(err)
	}

	formatter := &Formatter{
		Template: template,
		BufferPool: &sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
	}

	var testCases = []struct {
		name   string
		record logger.Record
	}{
		{
			"empty fields",
			logger.Record{
				Timestamp: time.Now(),
				Caller:    logger.Caller(2),
				Level:     logger.DEBUG,
				Message:   "test message",
				Fields:    make([]log.Field, 0),
			},
		},
		{
			"one field",
			logger.Record{
				Timestamp: time.Now(),
				Caller:    logger.Caller(2),
				Level:     logger.DEBUG,
				Message:   "test message",
				Fields: log.Fields(map[string]interface{}{
					"field": "test",
				}),
			},
		},
		{
			"several fields",
			logger.Record{
				Timestamp: time.Now(),
				Caller:    logger.Caller(2),
				Level:     logger.DEBUG,
				Message:   "test message",
				Fields: log.Fields(map[string]interface{}{
					"field_bool":   true,
					"field_int":    1,
					"field_float":  1.0,
					"field_string": "test",
					"field_error":  errors.New("test error"),
					"field_any":    time.Now(),
				}),
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			_, err = formatter.Format(testCase.record)
			if err != nil {
				t.Error(err)
			}
		})
	}
}
