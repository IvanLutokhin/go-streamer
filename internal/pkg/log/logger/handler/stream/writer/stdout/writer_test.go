package stdout

import (
	"github.com/IvanLutokhin/go-streamer/internal/pkg/log/logger/handler"
	"testing"
)

func TestNewWriter(t *testing.T) {
	_, err := NewWriter(handler.Options{})
	if err != nil {
		t.Error(err)
	}
}

func TestValidateWriter(t *testing.T) {
	err := ValidateWriter(handler.Options{})
	if err != nil {
		t.Error(err)
	}
}
