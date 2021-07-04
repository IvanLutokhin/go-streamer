package logger

import (
	"github.com/IvanLutokhin/go-streamer/internal/pkg/logger"
	"testing"
)

func TestRegisterHandlerWithEmptyCode(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic")
		}
	}()

	RegisterHandler(
		"",
		func(options map[string]interface{}) (logger.Handler, error) { return nil, nil },
		func(options map[string]interface{}) error { return nil },
	)
}

func TestRegisterHandlerWithNilNewHandlerFunc(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic")
		}
	}()

	RegisterHandler(
		"test",
		nil,
		func(options map[string]interface{}) error { return nil },
	)
}

func TestRegisterHandlerWithNilValidateHandlerFunc(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic")
		}
	}()

	RegisterHandler(
		"test",
		func(options map[string]interface{}) (logger.Handler, error) { return nil, nil },
		nil,
	)
}

func TestRegisterHandlerWithExistCode(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic")
		}
	}()

	t.Cleanup(func() {
		ResetHandlers()
	})

	RegisterHandler(
		"test",
		func(options map[string]interface{}) (logger.Handler, error) { return nil, nil },
		func(options map[string]interface{}) error { return nil },
	)

	RegisterHandler(
		"test",
		func(options map[string]interface{}) (logger.Handler, error) { return nil, nil },
		func(options map[string]interface{}) error { return nil },
	)
}

func TestNewHandler(t *testing.T) {
	t.Cleanup(func() {
		ResetHandlers()
	})

	RegisterHandler(
		"test",
		func(options map[string]interface{}) (logger.Handler, error) { return nil, nil },
		func(options map[string]interface{}) error { return nil },
	)

	_, err := NewHandler("test", nil)
	if err != nil {
		t.Error(err)
	}
}

func TestNewHandlerWithUnknownCode(t *testing.T) {
	_, err := NewHandler("test", nil)
	if err == nil {
		t.Error("expected error")
	}
}

func TestValidateHandler(t *testing.T) {
	t.Cleanup(func() {
		ResetHandlers()
	})

	RegisterHandler(
		"test",
		func(options map[string]interface{}) (logger.Handler, error) { return nil, nil },
		func(options map[string]interface{}) error { return nil },
	)

	err := ValidateHandler("test", nil)
	if err != nil {
		t.Error(err)
	}
}

func TestValidateHandlerWithUnknownCode(t *testing.T) {
	err := ValidateHandler("test", nil)
	if err == nil {
		t.Error("expected error")
	}
}
