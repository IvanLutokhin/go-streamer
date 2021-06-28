package text

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"
)

func TestParseTemplate(t *testing.T) {
	var invalidTestCases = []struct {
		name     string
		template string
	}{
		{"empty template", ""},
		{"template without tags", "test template"},
		{"invalid template", "test template with %invalid_tag"},
	}

	for _, invalidTestCase := range invalidTestCases {
		t.Run(invalidTestCase.name, func(t *testing.T) {
			_, err := ParseTemplate(invalidTestCase.template, true)
			if err == nil {
				t.Error("expected error")
			}
		})
	}

	t.Run("valid template", func(t *testing.T) {
		_, err := ParseTemplate("test template with %valid_tag%", true)
		if err != nil {
			t.Error(err)
		}
	})
}

func TestTemplate_Execute(t *testing.T) {
	var testCases = []struct {
		name           string
		template       string
		skipUnknownTag bool
		expectedString string
	}{
		{
			"skip unknown tag/not exist unknown tag",
			"test template %test_tag%",
			true,
			"test template success",
		},
		{
			"skip unknown tag/exist unknown tag",
			"test template %test_tag% %unknown_tag%",
			true,
			"test template success ",
		},
		{
			"not skip unknown tag/not exist unknown tag",
			"test template %test_tag%",
			false,
			"test template success",
		},
		{
			"not skip unknown tag/exist unknown tag",
			"test template %test_tag% %unknown_tag%",
			false,
			"test template success %unknown_tag%",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			template, err := ParseTemplate(testCase.template, testCase.skipUnknownTag)
			if err != nil {
				t.Error(err)
			}

			var buffer bytes.Buffer
			if err = template.Execute(&buffer, map[string]interface{}{"test_tag": "success"}); err != nil {
				t.Error(err)
			}

			if !strings.EqualFold(testCase.expectedString, buffer.String()) {
				t.Errorf("unexpected template")
			}
		})
	}
}

func TestTemplate_ExecuteWithValidValues(t *testing.T) {
	var testCases = []struct {
		name           string
		value          interface{}
		expectedString string
	}{
		{
			"bytes",
			[]byte{116, 101, 115, 116},
			"template test",
		},
		{
			"string",
			"test",
			"template test",
		},
		{
			"func",
			TagFunc(func(w io.Writer, tag string) error {
				if _, err := w.Write([]byte("test")); err != nil {
					return err
				}

				return nil
			}),
			"template test",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			template, err := ParseTemplate("template %test_tag%", true)
			if err != nil {
				t.Error(err)
			}

			var buffer bytes.Buffer
			if err = template.Execute(&buffer, map[string]interface{}{"test_tag": testCase.value}); err != nil {
				t.Error(err)
			}

			if !strings.EqualFold(testCase.expectedString, buffer.String()) {
				t.Errorf("unexpected template")
			}
		})
	}
}

func TestTemplate_ExecuteWithInvalidValues(t *testing.T) {
	var testCases = []struct {
		name  string
		value interface{}
	}{
		{
			"int",
			5,
		},
		{
			"invalid func",
			TagFunc(func(w io.Writer, tag string) error {
				return errors.New("test error")
			}),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			template, err := ParseTemplate("template %test_tag%", true)
			if err != nil {
				t.Error(err)
			}

			var buffer bytes.Buffer
			if err = template.Execute(&buffer, map[string]interface{}{"test_tag": testCase.value}); err == nil {
				t.Error("expected error")
			}
		})
	}
}
