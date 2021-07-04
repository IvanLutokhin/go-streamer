package stdout

import (
	"testing"
)

func TestNewWriter(t *testing.T) {
	_, err := NewWriter(nil)
	if err != nil {
		t.Error(err)
	}
}

func TestValidateWriter(t *testing.T) {
	err := ValidateWriter(map[string]interface{}{})
	if err != nil {
		t.Error(err)
	}
}
