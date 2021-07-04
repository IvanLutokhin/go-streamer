package stream

import (
	"github.com/IvanLutokhin/go-streamer/internal/pkg/logger/handlers/stream"
	"testing"
)

func TestRegisterFormatterWithEmptyCode(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic")
		}
	}()

	RegisterFormatter(
		"",
		func(options map[string]interface{}) (stream.Formatter, error) { return nil, nil },
		func(options map[string]interface{}) error { return nil },
	)
}

func TestRegisterFormatterWithNilNewFormatterFunc(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic")
		}
	}()

	RegisterFormatter(
		"test",
		nil,
		func(options map[string]interface{}) error { return nil },
	)
}

func TestRegisterFormatterWithNilValidateFormatterFunc(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic")
		}
	}()

	RegisterFormatter(
		"test",
		func(options map[string]interface{}) (stream.Formatter, error) { return nil, nil },
		nil,
	)
}

func TestRegisterFormatterWithExistCode(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic")
		}
	}()

	t.Cleanup(func() {
		ResetFormatters()
	})

	RegisterFormatter(
		"test",
		func(options map[string]interface{}) (stream.Formatter, error) { return nil, nil },
		func(options map[string]interface{}) error { return nil },
	)

	RegisterFormatter(
		"test",
		func(options map[string]interface{}) (stream.Formatter, error) { return nil, nil },
		func(options map[string]interface{}) error { return nil },
	)
}

func TestNewFormatter(t *testing.T) {
	t.Cleanup(func() {
		ResetFormatters()
	})

	RegisterFormatter(
		"test",
		func(options map[string]interface{}) (stream.Formatter, error) { return nil, nil },
		func(options map[string]interface{}) error { return nil },
	)

	_, err := NewFormatter("test", map[string]interface{}{})
	if err != nil {
		t.Error(err)
	}
}

func TestNewFormatterWithUnknownCode(t *testing.T) {
	_, err := NewFormatter("test", nil)
	if err == nil {
		t.Error("expected error")
	}
}

func TestValidateFormatter(t *testing.T) {
	t.Cleanup(func() {
		ResetFormatters()
	})

	RegisterFormatter(
		"test",
		func(options map[string]interface{}) (stream.Formatter, error) { return nil, nil },
		func(options map[string]interface{}) error { return nil },
	)

	err := ValidateFormatter("test", map[string]interface{}{})
	if err != nil {
		t.Error(err)
	}
}

func TestValidateFormatterWithUnknownCode(t *testing.T) {
	err := ValidateFormatter("test", map[string]interface{}{})
	if err == nil {
		t.Error("expected error")
	}
}
