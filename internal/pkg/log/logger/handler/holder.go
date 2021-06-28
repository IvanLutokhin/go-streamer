package handler

import (
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
)

type holder struct {
	code                string
	newHandlerFunc      func(options Options) (Handler, error)
	validateHandlerFunc func(options Options) error
}

var (
	holdersMu     sync.Mutex
	atomicHolders atomic.Value
)

func Reset() {
	atomicHolders.Store(make([]holder, 0))
}

func RegisterHandler(code string, newHandlerFunc func(options Options) (Handler, error), validateHandlerFunc func(options Options) error) {
	if strings.TrimSpace(code) == "" {
		panic("logger: handler code is empty")
	}

	if newHandlerFunc == nil {
		panic(fmt.Errorf("logger: constructor for handler %q must be specified", code))
	}

	if validateHandlerFunc == nil {
		panic(fmt.Errorf("logger: validator for handler %q must be specified", code))
	}

	holdersMu.Lock()
	defer holdersMu.Unlock()

	holders, _ := atomicHolders.Load().([]holder)
	for _, holder := range holders {
		if holder.code == code {
			panic(fmt.Errorf("logger: handler %q already exist", code))
		}
	}

	atomicHolders.Store(append(holders, holder{code, newHandlerFunc, validateHandlerFunc}))
}

func NewHandler(code string, options Options) (Handler, error) {
	holders, _ := atomicHolders.Load().([]holder)
	for _, holder := range holders {
		if holder.code == code {
			return holder.newHandlerFunc(options)
		}
	}

	return nil, fmt.Errorf("logger: handler %q was not registered", code)
}

func ValidateHandler(code string, options Options) error {
	holders, _ := atomicHolders.Load().([]holder)
	for _, holder := range holders {
		if holder.code == code {
			return holder.validateHandlerFunc(options)
		}
	}

	return fmt.Errorf("logger: handler %q was not registered", code)
}
