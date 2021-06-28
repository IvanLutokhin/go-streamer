package logger

import (
	"errors"
	"github.com/IvanLutokhin/go-streamer/internal/pkg/log"
	"github.com/IvanLutokhin/go-streamer/internal/pkg/log/logger/handler"
	"testing"
	"time"
)

type DummyHandler struct {
	records int
}

func (handler *DummyHandler) IsHandling(record handler.Record) bool {
	return true
}

func (handler *DummyHandler) Handle(record handler.Record) error {
	handler.records++

	return nil
}

func (handler *DummyHandler) Reset() {
	handler.records = 0
}

func TestLogger(t *testing.T) {
	t.Cleanup(func() {
		handler.Reset()
	})

	handler.Reset()
	handler.RegisterHandler(
		"test",
		func(options handler.Options) (handler.Handler, error) { return &DummyHandler{}, nil },
		func(options handler.Options) error { return nil },
	)

	h, err := handler.NewHandler("test", handler.Options{})
	if err != nil {
		t.Error(err)
	}

	logger := New(h)

	var testCases = []struct {
		name     string
		testFunc func(message string, fields ...log.Field)
	}{
		{"EMERGENCY", logger.Emergency},
		{"ALERT", logger.Alert},
		{"CRITICAL", logger.Critical},
		{"ERROR", logger.Error},
		{"WARNING", logger.Warning},
		{"NOTICE", logger.Notice},
		{"INFO", logger.Info},
		{"DEBUG", logger.Debug},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.testFunc("test message")
			testCase.testFunc("test message", log.FieldBool("test_field_bool", true))
			testCase.testFunc(
				"test message",
				log.FieldBool("test_field_bool", false),
				log.FieldInt("test_field_int", 123456),
				log.FieldFloat64("test_field_float64", 123.456),
				log.FieldString("test_field_string", "value"),
				log.FieldError("test_field_error", errors.New("test error")),
				log.FieldAny("test_field_any", time.Now()),
			)
			testCase.testFunc("test message", log.Fields(map[string]interface{}{"key": "value"})...)
			testCase.testFunc("test message", log.Fields(map[string]interface{}{"key": "value", "test": "test"})...)

			if h.(*DummyHandler).records != 5 {
				t.Error("unexpected record count")
			}

			h.(*DummyHandler).Reset()
		})
	}
}
