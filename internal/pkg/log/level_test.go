package log

import (
	"strings"
	"testing"
)

func TestLevelFromString(t *testing.T) {
	var testCases = []struct {
		name     string
		actual   string
		expected Level
	}{
		{"EMERGENCY/upper case", "EMERGENCY", EMERGENCY},
		{"EMERGENCY/lower case", "emergency", EMERGENCY},
		{"ALERT/upper case", "ALERT", ALERT},
		{"ALERT/lower case", "alert", ALERT},
		{"CRITICAL/upper case", "CRITICAL", CRITICAL},
		{"CRITICAL/lower case", "critical", CRITICAL},
		{"ERROR/upper case", "ERROR", ERROR},
		{"ERROR/lower case", "error", ERROR},
		{"WARNING/upper case", "WARNING", WARNING},
		{"WARNING/lower case", "warning", WARNING},
		{"NOTICE/upper case", "NOTICE", NOTICE},
		{"NOTICE/lower case", "notice", NOTICE},
		{"INFO/upper case", "INFO", INFO},
		{"INFO/lower case", "info", INFO},
		{"DEBUG/upper case", "DEBUG", DEBUG},
		{"DEBUG/lower case", "debug", DEBUG},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			level, err := LevelFromString(testCase.actual)
			if err != nil {
				t.Error(err)
			}

			if level != testCase.expected {
				t.Errorf("LevelFromString(%q) = %q; want %q", testCase.actual, level, testCase.expected)
			}
		})
	}

	t.Run("UNKNOWN", func(t *testing.T) {
		level, err := LevelFromString("invalid")
		if err == nil {
			t.Error("expected error")
		}

		if level != UNKNOWN {
			t.Errorf("LevelFromString('invalid') = %q; want %q", level, UNKNOWN)
		}
	})
}

func TestLevel_String(t *testing.T) {
	var testCases = []struct {
		name     string
		actual   Level
		expected string
	}{
		{"EMERGENCY", EMERGENCY, "EMERGENCY"},
		{"ALERT", ALERT, "ALERT"},
		{"CRITICAL", CRITICAL, "CRITICAL"},
		{"ERROR", ERROR, "ERROR"},
		{"WARNING", WARNING, "WARNING"},
		{"NOTICE", NOTICE, "NOTICE"},
		{"INFO", INFO, "INFO"},
		{"DEBUG", DEBUG, "DEBUG"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if !strings.EqualFold(testCase.actual.String(), testCase.expected) {
				t.Errorf("String() = %q; want %q", testCase.actual.String(), testCase.expected)
			}
		})
	}

	t.Run("INVALID", func(t *testing.T) {
		if !strings.EqualFold("Level(99)", Level(99).String()) {
			t.Error("unexpected level string")
		}
	})
}
