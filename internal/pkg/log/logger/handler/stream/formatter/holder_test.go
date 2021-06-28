package formatter

import (
	"github.com/IvanLutokhin/go-streamer/internal/pkg/log/logger/handler"
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
		func(options handler.Options) (Formatter, error) { return nil, nil },
		func(options handler.Options) error { return nil },
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
		func(options handler.Options) error { return nil },
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
		func(options handler.Options) (Formatter, error) { return nil, nil },
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
		Reset()
	})

	RegisterFormatter(
		"test",
		func(options handler.Options) (Formatter, error) { return nil, nil },
		func(options handler.Options) error { return nil },
	)

	RegisterFormatter(
		"test",
		func(options handler.Options) (Formatter, error) { return nil, nil },
		func(options handler.Options) error { return nil },
	)
}

func TestNewFormatter(t *testing.T) {
	t.Cleanup(func() {
		Reset()
	})

	RegisterFormatter(
		"test",
		func(options handler.Options) (Formatter, error) { return nil, nil },
		func(options handler.Options) error { return nil },
	)

	_, err := NewFormatter("test", handler.Options{})
	if err != nil {
		t.Error(err)
	}
}

func TestNewFormatterWithUnknownCode(t *testing.T) {
	_, err := NewFormatter("test", handler.Options{})
	if err == nil {
		t.Error("expected error")
	}
}

func TestValidateFormatter(t *testing.T) {
	t.Cleanup(func() {
		Reset()
	})

	RegisterFormatter(
		"test",
		func(options handler.Options) (Formatter, error) { return nil, nil },
		func(options handler.Options) error { return nil },
	)

	err := ValidateFormatter("test", handler.Options{})
	if err != nil {
		t.Error(err)
	}
}

func TestValidateFormatterWithUnknownCode(t *testing.T) {
	err := ValidateFormatter("test", handler.Options{})
	if err == nil {
		t.Error("expected error")
	}
}
