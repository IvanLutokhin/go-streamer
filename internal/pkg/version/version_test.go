package version

import (
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	title = "Test"
	tag = "v1.0.0"
	commit = "000000"

	s := String()
	if !strings.EqualFold(s, "Test v1.0.0 (Build: 000000)") {
		t.Errorf("unexpected version %q", s)
	}
}
