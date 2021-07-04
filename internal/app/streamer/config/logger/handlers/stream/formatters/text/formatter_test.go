package text

import (
	"testing"
)

func TestNewFormatterWithDefaultOptions(t *testing.T) {
	_, err := NewFormatter(nil)
	if err != nil {
		t.Error(err)
	}
}

func TestNewFormatterWithInvalidOptions(t *testing.T) {
	_, err := NewFormatter(map[string]interface{}{
		"template":       false,
		"skipUnknownTag": "yes",
	})
	if err == nil {
		t.Error("expected error")
	}
}

func TestNewFormatterWithEmptyTemplate(t *testing.T) {
	_, err := NewFormatter(map[string]interface{}{
		"template":       "",
		"skipUnknownTag": true,
	})
	if err == nil {
		t.Error("expected error")
	}
}

func TestValidateFormatterWithDefaultOptions(t *testing.T) {
	err := ValidateFormatter(map[string]interface{}{})
	if err != nil {
		t.Error(err)
	}
}

func TestValidateFormatterWithInvalidOptions(t *testing.T) {
	err := ValidateFormatter(map[string]interface{}{
		"template":       false,
		"skipUnknownTag": "yes",
	})
	if err == nil {
		t.Error("expected error")
	}
}

func TestValidateFormatterWithEmptyTemplate(t *testing.T) {
	err := ValidateFormatter(map[string]interface{}{
		"template":       "",
		"skipUnknownTag": true,
	})
	if err == nil {
		t.Error("expected error")
	}
}
