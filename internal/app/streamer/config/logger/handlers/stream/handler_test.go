package stream

import (
	"github.com/IvanLutokhin/go-streamer/internal/pkg/logger/handlers/stream"
	"io"
	"testing"
)

func TestNewHandlerWithDefaultOptions(t *testing.T) {
	t.Cleanup(func() {
		ResetFormatters()
		ResetWriters()
	})

	ResetFormatters()
	RegisterFormatter(
		"text",
		func(options map[string]interface{}) (stream.Formatter, error) { return nil, nil },
		func(options map[string]interface{}) error { return nil },
	)

	ResetWriters()
	RegisterWriter(
		"stdout",
		func(options map[string]interface{}) (io.Writer, error) { return nil, nil },
		func(options map[string]interface{}) error { return nil },
	)

	_, err := NewHandler(nil)
	if err != nil {
		t.Error(err)
	}
}

func TestNewHandlerWithInvalidOptions(t *testing.T) {
	_, err := NewHandler(map[string]interface{}{
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
		ResetFormatters()
		ResetWriters()
	})

	RegisterFormatter(
		"test_formatter",
		func(options map[string]interface{}) (stream.Formatter, error) { return nil, nil },
		func(options map[string]interface{}) error { return nil },
	)

	RegisterWriter(
		"test_writer",
		func(options map[string]interface{}) (io.Writer, error) { return nil, nil },
		func(options map[string]interface{}) error { return nil },
	)

	_, err := NewHandler(map[string]interface{}{
		"level": "invalid",
		"formatter": map[string]interface{}{
			"code": "test_formatter",
		},
		"writer": map[string]interface{}{
			"code": "test_writer",
		},
	})
	if err == nil {
		t.Error("expected error")
	}
}

func TestNewHandlerWithUnknownFormatter(t *testing.T) {
	t.Cleanup(func() {
		ResetWriters()
	})

	RegisterWriter(
		"test_writer",
		func(options map[string]interface{}) (io.Writer, error) { return nil, nil },
		func(options map[string]interface{}) error { return nil },
	)

	_, err := NewHandler(map[string]interface{}{
		"level": "debug",
		"formatter": map[string]interface{}{
			"code": "test_formatter",
		},
		"writer": map[string]interface{}{
			"code": "test_writer",
		},
	})
	if err == nil {
		t.Error("expected error")
	}
}

func TestNewHandlerWithUnknownWriter(t *testing.T) {
	t.Cleanup(func() {
		ResetFormatters()
	})

	RegisterFormatter(
		"test_formatter",
		func(options map[string]interface{}) (stream.Formatter, error) { return nil, nil },
		func(options map[string]interface{}) error { return nil },
	)

	_, err := NewHandler(map[string]interface{}{
		"level": "debug",
		"formatter": map[string]interface{}{
			"code": "test_formatter",
		},
		"writer": map[string]interface{}{
			"code": "test_writer",
		},
	})
	if err == nil {
		t.Error("expected error")
	}
}

func TestValidateHandlerWithDefaultOptions(t *testing.T) {
	t.Cleanup(func() {
		ResetFormatters()
		ResetWriters()
	})

	RegisterFormatter(
		"text",
		func(options map[string]interface{}) (stream.Formatter, error) { return nil, nil },
		func(options map[string]interface{}) error { return nil },
	)

	RegisterWriter(
		"stdout",
		func(options map[string]interface{}) (io.Writer, error) { return nil, nil },
		func(options map[string]interface{}) error { return nil },
	)

	err := ValidateHandler(map[string]interface{}{})
	if err != nil {
		t.Error(err)
	}
}

func TestValidateHandlerWithInvalidOptions(t *testing.T) {
	err := ValidateHandler(map[string]interface{}{
		"level": false,
	})
	if err == nil {
		t.Error("expected error")
	}
}

func TestValidateHandlerWithUnknownLevel(t *testing.T) {
	t.Cleanup(func() {
		ResetFormatters()
		ResetWriters()
	})

	RegisterFormatter(
		"test_formatter",
		func(options map[string]interface{}) (stream.Formatter, error) { return nil, nil },
		func(options map[string]interface{}) error { return nil },
	)

	RegisterWriter(
		"test_writer",
		func(options map[string]interface{}) (io.Writer, error) { return nil, nil },
		func(options map[string]interface{}) error { return nil },
	)

	err := ValidateHandler(map[string]interface{}{
		"level": "invalid",
		"formatter": map[string]interface{}{
			"code": "test_formatter",
		},
		"writer": map[string]interface{}{
			"code": "test_writer",
		},
	})
	if err == nil {
		t.Error("expected error")
	}
}

func TestValidateHandlerWithUnknownFormatter(t *testing.T) {
	t.Cleanup(func() {
		ResetWriters()
	})

	RegisterWriter(
		"test_writer",
		func(options map[string]interface{}) (io.Writer, error) { return nil, nil },
		func(options map[string]interface{}) error { return nil },
	)

	err := ValidateHandler(map[string]interface{}{
		"level": "debug",
		"formatter": map[string]interface{}{
			"code": "test_formatter",
		},
		"writer": map[string]interface{}{
			"code": "test_writer",
		},
	})
	if err == nil {
		t.Error("expected error")
	}
}

func TestValidateHandlerWithUnknownWriter(t *testing.T) {
	t.Cleanup(func() {
		ResetFormatters()
	})

	RegisterFormatter(
		"test_formatter",
		func(options map[string]interface{}) (stream.Formatter, error) { return nil, nil },
		func(options map[string]interface{}) error { return nil },
	)

	err := ValidateHandler(map[string]interface{}{
		"level": "debug",
		"formatter": map[string]interface{}{
			"code": "test_formatter",
		},
		"writer": map[string]interface{}{
			"code": "test_writer",
		},
	})
	if err == nil {
		t.Error("expected error")
	}
}
