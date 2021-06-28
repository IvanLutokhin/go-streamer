package formatter

import (
	"fmt"
	"github.com/IvanLutokhin/go-streamer/internal/pkg/log/logger/handler"
	"strings"
	"sync"
	"sync/atomic"
)

type holder struct {
	code                  string
	newFormatterFunc      func(options handler.Options) (Formatter, error)
	validateFormatterFunc func(options handler.Options) error
}

var (
	holdersMu     sync.Mutex
	atomicHolders atomic.Value
)

func Reset() {
	atomicHolders.Store(make([]holder, 0))
}

func RegisterFormatter(code string, newFormatterFunc func(options handler.Options) (Formatter, error), validateFormatterFunc func(options handler.Options) error) {
	if strings.TrimSpace(code) == "" {
		panic("logger: formatter code is empty")
	}

	if newFormatterFunc == nil {
		panic(fmt.Errorf("logger: constructor for formatter %q must be specified", code))
	}

	if validateFormatterFunc == nil {
		panic(fmt.Errorf("logger: validator for formatter %q must be specified", code))
	}

	holdersMu.Lock()
	defer holdersMu.Unlock()

	holders, _ := atomicHolders.Load().([]holder)
	for _, holder := range holders {
		if holder.code == code {
			panic(fmt.Errorf("logger: formatter %q already exist", code))
		}
	}

	atomicHolders.Store(append(holders, holder{code, newFormatterFunc, validateFormatterFunc}))
}

func NewFormatter(code string, options handler.Options) (Formatter, error) {
	holders, _ := atomicHolders.Load().([]holder)
	for _, holder := range holders {
		if holder.code == code {
			return holder.newFormatterFunc(options)
		}
	}

	return nil, fmt.Errorf("logger: formatter %q was not registered", code)
}

func ValidateFormatter(code string, options handler.Options) error {
	holders, _ := atomicHolders.Load().([]holder)
	for _, holder := range holders {
		if holder.code == code {
			return holder.validateFormatterFunc(options)
		}
	}

	return fmt.Errorf("logger: formatter %q was not registered", code)
}
