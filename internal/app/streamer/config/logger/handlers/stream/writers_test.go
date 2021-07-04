package stream

import (
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
		func(options map[string]interface{}) (io.Writer, error) { return nil, nil },
		func(options map[string]interface{}) error { return nil },
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
		func(options map[string]interface{}) error { return nil },
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
		func(options map[string]interface{}) (io.Writer, error) { return nil, nil },
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
		ResetWriters()
	})

	RegisterWriter(
		"test",
		func(options map[string]interface{}) (io.Writer, error) { return nil, nil },
		func(options map[string]interface{}) error { return nil },
	)

	RegisterWriter(
		"test",
		func(options map[string]interface{}) (io.Writer, error) { return nil, nil },
		func(options map[string]interface{}) error { return nil },
	)
}

func TestNewWriter(t *testing.T) {
	t.Cleanup(func() {
		ResetWriters()
	})

	RegisterWriter(
		"test",
		func(options map[string]interface{}) (io.Writer, error) { return nil, nil },
		func(options map[string]interface{}) error { return nil },
	)

	_, err := NewWriter("test", map[string]interface{}{})
	if err != nil {
		t.Error(err)
	}
}

func TestNewWriterWithUnknownCode(t *testing.T) {
	_, err := NewWriter("test", nil)
	if err == nil {
		t.Error("expected error")
	}
}

func TestValidateWriter(t *testing.T) {
	t.Cleanup(func() {
		ResetWriters()
	})

	RegisterWriter(
		"test",
		func(options map[string]interface{}) (io.Writer, error) { return nil, nil },
		func(options map[string]interface{}) error { return nil },
	)

	err := ValidateWriter("test", map[string]interface{}{})
	if err != nil {
		t.Error(err)
	}
}

func TestValidateWriterWithUnknownCode(t *testing.T) {
	err := ValidateWriter("test", map[string]interface{}{})
	if err == nil {
		t.Error("expected error")
	}
}
