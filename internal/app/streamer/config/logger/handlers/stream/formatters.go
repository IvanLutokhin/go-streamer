package stream

import (
	"fmt"
	"github.com/IvanLutokhin/go-streamer/internal/pkg/logger/handlers/stream"
	"strings"
	"sync"
	"sync/atomic"
)

type formatter struct {
	code                  string
	newFormatterFunc      func(options map[string]interface{}) (stream.Formatter, error)
	validateFormatterFunc func(options map[string]interface{}) error
}

var (
	formattersMu     sync.Mutex
	atomicFormatters atomic.Value
)

func ResetFormatters() {
	atomicFormatters.Store(make([]formatter, 0))
}

func RegisterFormatter(code string, newFormatterFunc func(options map[string]interface{}) (stream.Formatter, error), validateFormatterFunc func(options map[string]interface{}) error) {
	if strings.TrimSpace(code) == "" {
		panic("config: logger: formatter code is empty")
	}

	if newFormatterFunc == nil {
		panic(fmt.Errorf("config: logger: constructor for formatter %q must be specified", code))
	}

	if validateFormatterFunc == nil {
		panic(fmt.Errorf("config: logger: validator for formatter %q must be specified", code))
	}

	formattersMu.Lock()
	defer formattersMu.Unlock()

	formatters, _ := atomicFormatters.Load().([]formatter)
	for _, formatter := range formatters {
		if formatter.code == code {
			panic(fmt.Errorf("config: logger: formatter %q already exist", code))
		}
	}

	atomicFormatters.Store(append(formatters, formatter{code, newFormatterFunc, validateFormatterFunc}))
}

func NewFormatter(code string, options map[string]interface{}) (stream.Formatter, error) {
	formatters, _ := atomicFormatters.Load().([]formatter)
	for _, formatter := range formatters {
		if formatter.code == code {
			return formatter.newFormatterFunc(options)
		}
	}

	return nil, fmt.Errorf("config: logger: formatter %q was not registered", code)
}

func ValidateFormatter(code string, options map[string]interface{}) error {
	formatters, _ := atomicFormatters.Load().([]formatter)
	for _, formatter := range formatters {
		if formatter.code == code {
			return formatter.validateFormatterFunc(options)
		}
	}

	return fmt.Errorf("config: logger: formatter %q was not registered", code)
}
