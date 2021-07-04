package stream

import (
	"bytes"
	"errors"
	"github.com/IvanLutokhin/go-streamer/internal/pkg/logger"
	"github.com/IvanLutokhin/go-streamer/internal/pkg/logger/handlers/stream/formatters/text"
	"github.com/IvanLutokhin/go-streamer/pkg/log"
	"os"
	"sync"
	"testing"
)

func BenchmarkStreamHandlerWithTextFormatterSeparateFields(b *testing.B) {
	template, err := text.ParseTemplate("%datetime% %caller% %level% %message% %context%", false)
	if err != nil {
		b.Error(err)
	}

	textFormatter := &text.Formatter{
		Template: template,
		BufferPool: &sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
	}

	devNullWriter, _ := os.OpenFile(os.DevNull, os.O_WRONLY, 0666)

	defer devNullWriter.Close()

	stdLogger := logger.New(&Handler{Level: logger.DEBUG, Formatter: textFormatter, Writer: devNullWriter})

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			stdLogger.Debug(
				"test message",
				log.FieldBool("test_field_bool", true),
				log.FieldInt("test_field_int", 123456),
				log.FieldFloat64("test_field_float64", 123.456),
				log.FieldString("test_field_string", "value"),
				log.FieldError("test_field_error", errors.New("test error")),
				log.FieldAny("test_field_any", struct{}{}),
			)
		}
	})
}

func BenchmarkStreamHandlerWithTextFormatterCombinedFields(b *testing.B) {
	template, err := text.ParseTemplate("%datetime% %caller% %level% %message% %context%", false)
	if err != nil {
		b.Error(err)
	}

	textFormatter := &text.Formatter{
		Template: template,
		BufferPool: &sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
	}

	devNullWriter, _ := os.OpenFile(os.DevNull, os.O_WRONLY, 0666)

	defer devNullWriter.Close()

	stdLogger := logger.New(&Handler{Level: logger.DEBUG, Formatter: textFormatter, Writer: devNullWriter})

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			stdLogger.Debug(
				"test message",
				log.Fields(map[string]interface{}{
					"test_field_bool":    true,
					"test_field_int":     123456,
					"test_field_float64": 123.456,
					"test_field_string":  "value",
					"test_field_error":   errors.New("test error"),
					"test_field_any":     struct{}{},
				})...,
			)
		}
	})
}
