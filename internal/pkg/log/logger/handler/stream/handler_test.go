package stream

import (
	"errors"
	"github.com/IvanLutokhin/go-streamer/internal/pkg/log"
	"github.com/IvanLutokhin/go-streamer/internal/pkg/log/logger/handler"
	"github.com/IvanLutokhin/go-streamer/internal/pkg/log/logger/handler/stream/formatter"
	"github.com/IvanLutokhin/go-streamer/internal/pkg/log/logger/handler/stream/writer"
	"io"
	"testing"
)

type DummyFormatter struct {
	ReturnError bool
}

func (formatter *DummyFormatter) Format(record handler.Record) (bytes []byte, err error) {
	if !formatter.ReturnError {
		return
	}

	err = errors.New("format error")

	return
}

type DummyWriter struct {
	ReturnError bool
}

func (writer *DummyWriter) Write([]byte) (n int, err error) {
	if !writer.ReturnError {
		return
	}

	err = errors.New("write error")

	return
}

func TestNewHandlerWithDefaultOptions(t *testing.T) {
	t.Cleanup(func() {
		formatter.Reset()
		writer.Reset()
	})

	formatter.Reset()
	formatter.RegisterFormatter(
		"text",
		func(options handler.Options) (formatter.Formatter, error) { return nil, nil },
		func(options handler.Options) error { return nil },
	)

	writer.Reset()
	writer.RegisterWriter(
		"stdout",
		func(options handler.Options) (io.Writer, error) { return nil, nil },
		func(options handler.Options) error { return nil },
	)

	_, err := NewHandler(nil)
	if err != nil {
		t.Error(err)
	}
}

func TestNewHandlerWithInvalidOptions(t *testing.T) {
	_, err := NewHandler(handler.Options{
		"level":     false,
		"formatter": nil,
		"writer":    nil,
	})
	if err == nil {
		t.Error("expected error")
	}
}

func TestNewHandlerWithUnknownLevel(t *testing.T) {
	t.Cleanup(func() {
		formatter.Reset()
		writer.Reset()
	})

	formatter.RegisterFormatter(
		"test_formatter",
		func(options handler.Options) (formatter.Formatter, error) { return nil, nil },
		func(options handler.Options) error { return nil },
	)

	writer.RegisterWriter(
		"test_writer",
		func(options handler.Options) (io.Writer, error) { return nil, nil },
		func(options handler.Options) error { return nil },
	)

	_, err := NewHandler(handler.Options{
		"level": "invalid",
		"formatter": handler.Options{
			"code": "test_formatter",
		},
		"writer": handler.Options{
			"code": "test_writer",
		},
	})
	if err == nil {
		t.Error("expected error")
	}
}

func TestNewHandlerWithUnknownFormatter(t *testing.T) {
	t.Cleanup(func() {
		writer.Reset()
	})

	writer.RegisterWriter(
		"test_writer",
		func(options handler.Options) (io.Writer, error) { return nil, nil },
		func(options handler.Options) error { return nil },
	)

	_, err := NewHandler(handler.Options{
		"level": "debug",
		"formatter": handler.Options{
			"code": "test_formatter",
		},
		"writer": handler.Options{
			"code": "test_writer",
		},
	})
	if err == nil {
		t.Error("expected error")
	}
}

func TestNewHandlerWithUnknownWriter(t *testing.T) {
	t.Cleanup(func() {
		formatter.Reset()
	})

	formatter.RegisterFormatter(
		"test_formatter",
		func(options handler.Options) (formatter.Formatter, error) { return nil, nil },
		func(options handler.Options) error { return nil },
	)

	_, err := NewHandler(handler.Options{
		"level": "debug",
		"formatter": handler.Options{
			"code": "test_formatter",
		},
		"writer": handler.Options{
			"code": "test_writer",
		},
	})
	if err == nil {
		t.Error("expected error")
	}
}

func TestValidateHandlerWithDefaultOptions(t *testing.T) {
	t.Cleanup(func() {
		formatter.Reset()
		writer.Reset()
	})

	formatter.RegisterFormatter(
		"text",
		func(options handler.Options) (formatter.Formatter, error) { return nil, nil },
		func(options handler.Options) error { return nil },
	)

	writer.RegisterWriter(
		"stdout",
		func(options handler.Options) (io.Writer, error) { return nil, nil },
		func(options handler.Options) error { return nil },
	)

	err := ValidateHandler(handler.Options{})
	if err != nil {
		t.Error(err)
	}
}

func TestValidateHandlerWithInvalidOptions(t *testing.T) {
	err := ValidateHandler(handler.Options{
		"level": false,
	})
	if err == nil {
		t.Error("expected error")
	}
}

func TestValidateHandlerWithUnknownLevel(t *testing.T) {
	t.Cleanup(func() {
		formatter.Reset()
		writer.Reset()
	})

	formatter.RegisterFormatter(
		"test_formatter",
		func(options handler.Options) (formatter.Formatter, error) { return nil, nil },
		func(options handler.Options) error { return nil },
	)

	writer.RegisterWriter(
		"test_writer",
		func(options handler.Options) (io.Writer, error) { return nil, nil },
		func(options handler.Options) error { return nil },
	)

	err := ValidateHandler(handler.Options{
		"level": "invalid",
		"formatter": handler.Options{
			"code": "test_formatter",
		},
		"writer": handler.Options{
			"code": "test_writer",
		},
	})
	if err == nil {
		t.Errorf("expected error")
	}
}

func TestValidateHandlerWithUnknownFormatter(t *testing.T) {
	t.Cleanup(func() {
		writer.Reset()
	})

	writer.RegisterWriter(
		"test_writer",
		func(options handler.Options) (io.Writer, error) { return nil, nil },
		func(options handler.Options) error { return nil },
	)

	err := ValidateHandler(handler.Options{
		"level": "debug",
		"formatter": handler.Options{
			"code": "test_formatter",
		},
		"writer": handler.Options{
			"code": "test_writer",
		},
	})
	if err == nil {
		t.Errorf("expected error")
	}
}

func TestValidateHandlerWithUnknownWriter(t *testing.T) {
	t.Cleanup(func() {
		formatter.Reset()
	})

	formatter.RegisterFormatter(
		"test_formatter",
		func(options handler.Options) (formatter.Formatter, error) { return nil, nil },
		func(options handler.Options) error { return nil },
	)

	err := ValidateHandler(handler.Options{
		"level": "debug",
		"formatter": handler.Options{
			"code": "test_formatter",
		},
		"writer": handler.Options{
			"code": "test_writer",
		},
	})
	if err == nil {
		t.Errorf("expected error")
	}
}

func TestHandler_IsHandling(t *testing.T) {
	t.Cleanup(func() {
		formatter.Reset()
		writer.Reset()
	})

	formatter.RegisterFormatter(
		"test_formatter",
		func(options handler.Options) (formatter.Formatter, error) { return nil, nil },
		func(options handler.Options) error { return nil },
	)

	writer.RegisterWriter(
		"test_writer",
		func(options handler.Options) (io.Writer, error) { return nil, nil },
		func(options handler.Options) error { return nil },
	)

	h, err := NewHandler(handler.Options{
		"level": "info",
		"formatter": handler.Options{
			"code": "test_formatter",
		},
		"writer": handler.Options{
			"code": "test_writer",
		},
	})
	if err != nil {
		t.Error(err)
	}

	if ok := h.IsHandling(handler.Record{Level: log.DEBUG}); ok {
		t.Error("expected false")
	}

	if ok := h.IsHandling(handler.Record{Level: log.ERROR}); !ok {
		t.Errorf("expected true")
	}
}

func TestHandler_HandleWithErrorOnFormatting(t *testing.T) {
	t.Cleanup(func() {
		formatter.Reset()
		writer.Reset()
	})

	var testCases = []struct {
		name          string
		errorOnFormat bool
		errorOnWrite  bool
		expectError   bool
	}{
		{"format error/write not error", true, false, true},
		{"format not error/write error", false, true, true},
		{"format not error/write not error", false, false, false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			formatter.Reset()
			formatter.RegisterFormatter(
				"test_formatter",
				func(options handler.Options) (formatter.Formatter, error) {
					return &DummyFormatter{ReturnError: testCase.errorOnFormat}, nil
				},
				func(options handler.Options) error { return nil },
			)

			writer.Reset()
			writer.RegisterWriter(
				"test_writer",
				func(options handler.Options) (io.Writer, error) {
					return &DummyWriter{ReturnError: testCase.errorOnWrite}, nil
				},
				func(options handler.Options) error { return nil },
			)

			h, err := NewHandler(handler.Options{
				"level": "info",
				"formatter": handler.Options{
					"code": "test_formatter",
				},
				"writer": handler.Options{
					"code": "test_writer",
				},
			})
			if err != nil {
				t.Error(err)
			}

			err = h.Handle(handler.Record{})
			if err != nil {
				if !testCase.expectError {
					t.Error(err)
				}
			} else {
				if testCase.expectError {
					t.Error("expected error")
				}
			}
		})
	}
}
