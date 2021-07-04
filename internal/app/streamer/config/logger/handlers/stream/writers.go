package stream

import (
	"fmt"
	"io"
	"strings"
	"sync"
	"sync/atomic"
)

type writer struct {
	code               string
	newWriterFunc      func(options map[string]interface{}) (io.Writer, error)
	validateWriterFunc func(options map[string]interface{}) error
}

var (
	writersMu     sync.Mutex
	atomicWriters atomic.Value
)

func ResetWriters() {
	atomicWriters.Store(make([]writer, 0))
}

func RegisterWriter(code string, newWriterFunc func(options map[string]interface{}) (io.Writer, error), validateWriterFunc func(options map[string]interface{}) error) {
	if strings.TrimSpace(code) == "" {
		panic("config: logger: writer code is empty")
	}

	if newWriterFunc == nil {
		panic(fmt.Errorf("config: logger: constructor for writer %q must be specified", code))
	}

	if validateWriterFunc == nil {
		panic(fmt.Errorf("config: logger: validator for writer %q must be specified", code))
	}

	writersMu.Lock()
	defer writersMu.Unlock()

	writers, _ := atomicWriters.Load().([]writer)
	for _, writer := range writers {
		if writer.code == code {
			panic(fmt.Errorf("config: logger: writer %q already exist", code))
		}
	}

	atomicWriters.Store(append(writers, writer{code, newWriterFunc, validateWriterFunc}))
}

func NewWriter(code string, options map[string]interface{}) (io.Writer, error) {
	writers, _ := atomicWriters.Load().([]writer)
	for _, writer := range writers {
		if writer.code == code {
			return writer.newWriterFunc(options)
		}
	}

	return nil, fmt.Errorf("config: logger: writer %q was not registered", code)
}

func ValidateWriter(code string, options map[string]interface{}) error {
	writers, _ := atomicWriters.Load().([]writer)
	for _, writer := range writers {
		if writer.code == code {
			return writer.validateWriterFunc(options)
		}
	}

	return fmt.Errorf("config: logger: writer %q was not registered", code)
}
