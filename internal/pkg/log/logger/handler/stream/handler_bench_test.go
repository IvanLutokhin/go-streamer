package stream

import (
	"errors"
	"github.com/IvanLutokhin/go-streamer/internal/pkg/log"
	"github.com/IvanLutokhin/go-streamer/internal/pkg/log/logger"
	"github.com/IvanLutokhin/go-streamer/internal/pkg/log/logger/handler"
	"github.com/IvanLutokhin/go-streamer/internal/pkg/log/logger/handler/stream/formatter/text"
	"os"
	"testing"
)

func BenchmarkStreamHandlerWithTextFormatterSeparateFields(b *testing.B) {
	TextFormatter, _ := text.NewFormatter(handler.Options{})

	DevNullWriter, _ := os.OpenFile(os.DevNull, os.O_WRONLY, 0666)

	defer DevNullWriter.Close()

	StdLogger := logger.New(&Handler{level: log.DEBUG, formatter: TextFormatter, writer: DevNullWriter})

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			StdLogger.Debug(
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
	TextFormatter, _ := text.NewFormatter(handler.Options{})

	DevNullWriter, _ := os.OpenFile(os.DevNull, os.O_WRONLY, 0666)

	defer DevNullWriter.Close()

	StdLogger := logger.New(&Handler{level: log.DEBUG, formatter: TextFormatter, writer: DevNullWriter})

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			StdLogger.Debug(
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
