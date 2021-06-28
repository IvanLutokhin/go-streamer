package handler

import (
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
		func(options Options) (Handler, error) { return nil, nil },
		func(options Options) error { return nil },
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
		func(options Options) error { return nil },
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
		func(options Options) (Handler, error) { return nil, nil },
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
		Reset()
	})

	RegisterHandler(
		"test",
		func(options Options) (Handler, error) { return nil, nil },
		func(options Options) error { return nil },
	)

	RegisterHandler(
		"test",
		func(options Options) (Handler, error) { return nil, nil },
		func(options Options) error { return nil },
	)
}

func TestNewHandler(t *testing.T) {
	t.Cleanup(func() {
		Reset()
	})

	RegisterHandler(
		"test",
		func(options Options) (Handler, error) { return nil, nil },
		func(options Options) error { return nil },
	)

	_, err := NewHandler("test", Options{})
	if err != nil {
		t.Error(err)
	}
}

func TestNewHandlerWithUnknownCode(t *testing.T) {
	_, err := NewHandler("test", Options{})
	if err == nil {
		t.Error("expected error")
	}
}

func TestValidateHandler(t *testing.T) {
	t.Cleanup(func() {
		Reset()
	})

	RegisterHandler(
		"test",
		func(options Options) (Handler, error) { return nil, nil },
		func(options Options) error { return nil },
	)

	err := ValidateHandler("test", Options{})
	if err != nil {
		t.Error(err)
	}
}

func TestValidateHandlerWithUnknownCode(t *testing.T) {
	err := ValidateHandler("test", Options{})
	if err == nil {
		t.Error("expected error")
	}
}
