package stream

import (
	"errors"
	"github.com/IvanLutokhin/go-streamer/internal/pkg/logger"
	"testing"
)

type DummyFormatter struct {
	ReturnError bool
}

func (formatter *DummyFormatter) Format(logger.Record) (bytes []byte, err error) {
	if !formatter.ReturnError {
		return
	}

	err = errors.New("format error")

	return
}

type DummyWriter struct {
	ReturnError bool
}

func (writer *DummyWriter) Write([]byte) (n int, err error) {
	if !writer.ReturnError {
		return
	}

	err = errors.New("write error")

	return
}

func TestHandler_IsHandling(t *testing.T) {
	h := &Handler{
		Level:     logger.INFO,
		Formatter: &DummyFormatter{ReturnError: false},
		Writer:    &DummyWriter{ReturnError: false},
	}

	if ok := h.IsHandling(logger.Record{Level: logger.DEBUG}); ok {
		t.Error("expected false")
	}

	if ok := h.IsHandling(logger.Record{Level: logger.INFO}); !ok {
		t.Error("expected true")
	}

	if ok := h.IsHandling(logger.Record{Level: logger.ERROR}); !ok {
		t.Error("expected true")
	}
}

func TestHandler_HandleWithErrorOnFormatting(t *testing.T) {
	var testCases = []struct {
		name          string
		errorOnFormat bool
		errorOnWrite  bool
		expectError   bool
	}{
		{"format error/write not error", true, false, true},
		{"format not error/write error", false, true, true},
		{"format not error/write not error", false, false, false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			h := &Handler{
				Level:     logger.INFO,
				Formatter: &DummyFormatter{ReturnError: testCase.errorOnFormat},
				Writer:    &DummyWriter{ReturnError: testCase.errorOnWrite},
			}

			err := h.Handle(logger.Record{})
			if err != nil {
				if !testCase.expectError {
					t.Error(err)
				}
			} else {
				if testCase.expectError {
					t.Error("expected error")
				}
			}
		})
	}
}
