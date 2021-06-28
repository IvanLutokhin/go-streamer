package log

import (
	"errors"
	"testing"
	"time"
)

func TestFieldAny(t *testing.T) {
	var testCases = []struct {
		name         string
		testValue    interface{}
		expectedType FieldType
	}{
		{"bool", true, FieldBoolType},
		{"int", 1, FieldIntType},
		{"float64", 1.0, FieldFloat64Type},
		{"string", "test", FieldStringType},
		{"error", errors.New("test error"), FieldErrorType},
		{"unknown", time.Now(), FieldUnknownType},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			f := FieldAny("test_key", testCase.testValue)
			if f.Type != testCase.expectedType {
				t.Error("unexpected field type")
			}
		})
	}
}
