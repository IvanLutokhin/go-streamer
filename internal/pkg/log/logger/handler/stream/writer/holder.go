package writer

import (
	"fmt"
	"github.com/IvanLutokhin/go-streamer/internal/pkg/log/logger/handler"
	"io"
	"strings"
	"sync"
	"sync/atomic"
)

type holder struct {
	code               string
	newWriterFunc      func(options handler.Options) (io.Writer, error)
	validateWriterFunc func(options handler.Options) error
}

var (
	holdersMu     sync.Mutex
	atomicHolders atomic.Value
)

func Reset() {
	atomicHolders.Store(make([]holder, 0))
}

func RegisterWriter(code string, newWriterFunc func(options handler.Options) (io.Writer, error), validateWriterFunc func(options handler.Options) error) {
	if strings.TrimSpace(code) == "" {
		panic("logger: writer code is empty")
	}

	if newWriterFunc == nil {
		panic(fmt.Errorf("logger: constructor for writer %q must be specified", code))
	}

	if validateWriterFunc == nil {
		panic(fmt.Errorf("logger: validator for writer %q must be specified", code))
	}

	holdersMu.Lock()
	defer holdersMu.Unlock()

	holders, _ := atomicHolders.Load().([]holder)
	for _, holder := range holders {
		if holder.code == code {
			panic(fmt.Errorf("logger: writer %q already exist", code))
		}
	}

	atomicHolders.Store(append(holders, holder{code, newWriterFunc, validateWriterFunc}))
}

func NewWriter(code string, options handler.Options) (io.Writer, error) {
	holders, _ := atomicHolders.Load().([]holder)
	for _, holder := range holders {
		if holder.code == code {
			return holder.newWriterFunc(options)
		}
	}

	return nil, fmt.Errorf("logger: writer %q was not registered", code)
}

func ValidateWriter(code string, options handler.Options) error {
	holders, _ := atomicHolders.Load().([]holder)
	for _, holder := range holders {
		if holder.code == code {
			return holder.validateWriterFunc(options)
		}
	}

	return fmt.Errorf("logger: writer %q was not registered", code)
}
