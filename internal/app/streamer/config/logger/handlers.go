package logger

import (
	"fmt"
	"github.com/IvanLutokhin/go-streamer/internal/pkg/logger"
	"strings"
	"sync"
	"sync/atomic"
)

type handler struct {
	code                string
	newHandlerFunc      func(options map[string]interface{}) (logger.Handler, error)
	validateHandlerFunc func(options map[string]interface{}) error
}

var (
	handlersMu     sync.Mutex
	atomicHandlers atomic.Value
)

func ResetHandlers() {
	atomicHandlers.Store(make([]handler, 0))
}

func RegisterHandler(code string, newHandlerFunc func(options map[string]interface{}) (logger.Handler, error), validateHandlerFunc func(options map[string]interface{}) error) {
	if strings.TrimSpace(code) == "" {
		panic("config: logger: handler code is empty")
	}

	if newHandlerFunc == nil {
		panic(fmt.Errorf("config: logger: constructor for handler %q must be specified", code))
	}

	if validateHandlerFunc == nil {
		panic(fmt.Errorf("config: logger: validator for handler %q must be specified", code))
	}

	handlersMu.Lock()
	defer handlersMu.Unlock()

	handlers, _ := atomicHandlers.Load().([]handler)
	for _, handler := range handlers {
		if handler.code == code {
			panic(fmt.Errorf("config: logger: handler %q already exist", code))
		}
	}

	atomicHandlers.Store(append(handlers, handler{code, newHandlerFunc, validateHandlerFunc}))
}

func NewHandler(code string, options map[string]interface{}) (logger.Handler, error) {
	handlers, _ := atomicHandlers.Load().([]handler)
	for _, handler := range handlers {
		if handler.code == code {
			return handler.newHandlerFunc(options)
		}
	}

	return nil, fmt.Errorf("config: logger: handler %q was not registered", code)
}

func ValidateHandler(code string, options map[string]interface{}) error {
	handlers, _ := atomicHandlers.Load().([]handler)
	for _, handler := range handlers {
		if handler.code == code {
			return handler.validateHandlerFunc(options)
		}
	}

	return fmt.Errorf("config: logger: handler %q was not registered", code)
}
