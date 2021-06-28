package text

import (
	"errors"
	"github.com/IvanLutokhin/go-streamer/internal/pkg/log"
	"github.com/IvanLutokhin/go-streamer/internal/pkg/log/logger/handler"
	"testing"
	"time"
)

func TestNewFormatterWithDefaultOptions(t *testing.T) {
	_, err := NewFormatter(nil)
	if err != nil {
		t.Error(err)
	}
}

func TestNewFormatterWithInvalidOptions(t *testing.T) {
	_, err := NewFormatter(handler.Options{
		"template":       false,
		"skipUnknownTag": "yes",
	})
	if err == nil {
		t.Errorf("expected error")
	}
}

func TestNewFormatterWithInvalidTemplate(t *testing.T) {
	_, err := NewFormatter(handler.Options{
		"template":       "",
		"skipUnknownTag": true,
	})
	if err == nil {
		t.Errorf("expected error")
	}
}

func TestValidateFormatterWithDefaultOptions(t *testing.T) {
	err := ValidateFormatter(handler.Options{})
	if err != nil {
		t.Error(err)
	}
}

func TestValidateFormatterWithInvalidOptions(t *testing.T) {
	err := ValidateFormatter(handler.Options{
		"template":       false,
		"skipUnknownTag": "yes",
	})
	if err == nil {
		t.Error("expected error")
	}
}

func TestValidateFormatterWithEmptyTemplate(t *testing.T) {
	err := ValidateFormatter(handler.Options{
		"template":       "",
		"skipUnknownTag": true,
	})
	if err == nil {
		t.Error("expected error")
	}
}

func TestFormatter_Format(t *testing.T) {
	formatter, err := NewFormatter(handler.Options{})
	if err != nil {
		t.Error(err)
	}

	var testCases = []struct {
		name   string
		record handler.Record
	}{
		{
			"empty fields",
			handler.Record{
				Timestamp: time.Now(),
				Caller:    handler.Caller(2),
				Level:     log.DEBUG,
				Message:   "test message",
				Fields:    make([]log.Field, 0),
			},
		},
		{
			"one field",
			handler.Record{
				Timestamp: time.Now(),
				Caller:    handler.Caller(2),
				Level:     log.DEBUG,
				Message:   "test message",
				Fields: log.Fields(map[string]interface{}{
					"field": "test",
				}),
			},
		},
		{
			"several fields",
			handler.Record{
				Timestamp: time.Now(),
				Caller:    handler.Caller(2),
				Level:     log.DEBUG,
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
