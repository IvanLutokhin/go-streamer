package writer

import (
	"github.com/IvanLutokhin/go-streamer/internal/pkg/log/logger/handler"
	"io"
	"testing"
)

func TestRegisterWriterWithEmptyCode(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic")
		}
	}()

	RegisterWriter(
		"",
		func(options handler.Options) (io.Writer, error) { return nil, nil },
		func(options handler.Options) error { return nil },
	)
}

func TestRegisterWriterWithNilNewFormatterFunc(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic")
		}
	}()

	RegisterWriter(
		"test",
		nil,
		func(options handler.Options) error { return nil },
	)
}

func TestRegisterWriterWithNilValidateFormatterFunc(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic")
		}
	}()

	RegisterWriter(
		"test",
		func(options handler.Options) (io.Writer, error) { return nil, nil },
		nil,
	)
}

func TestRegisterWriterWithExistCode(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic")
		}
	}()

	t.Cleanup(func() {
		Reset()
	})

	RegisterWriter(
		"test",
		func(options handler.Options) (io.Writer, error) { return nil, nil },
		func(options handler.Options) error { return nil },
	)

	RegisterWriter(
		"test",
		func(options handler.Options) (io.Writer, error) { return nil, nil },
		func(options handler.Options) error { return nil },
	)
}

func TestNewWriter(t *testing.T) {
	t.Cleanup(func() {
		Reset()
	})

	RegisterWriter(
		"test",
		func(options handler.Options) (io.Writer, error) { return nil, nil },
		func(options handler.Options) error { return nil },
	)

	_, err := NewWriter("test", handler.Options{})
	if err != nil {
		t.Error(err)
	}
}

func TestNewWriterWithUnknownCode(t *testing.T) {
	_, err := NewWriter("test", handler.Options{})
	if err == nil {
		t.Error("expected error")
	}
}

func TestValidateWriter(t *testing.T) {
	t.Cleanup(func() {
		Reset()
	})

	RegisterWriter(
		"test",
		func(options handler.Options) (io.Writer, error) { return nil, nil },
		func(options handler.Options) error { return nil },
	)

	err := ValidateWriter("test", handler.Options{})
	if err != nil {
		t.Error(err)
	}
}

func TestValidateWriterWithUnknownCode(t *testing.T) {
	err := ValidateWriter("test", handler.Options{})
	if err == nil {
		t.Error("expected error")
	}
}
