package text

import (
	"bytes"
	"fmt"
	"github.com/IvanLutokhin/go-streamer/internal/pkg/logger"
	"github.com/IvanLutokhin/go-streamer/pkg/log"
	"io"
	"strings"
	"sync"
	"time"
)

type Formatter struct {
	Template   *Template
	BufferPool *sync.Pool
}

func (formatter *Formatter) Format(record logger.Record) ([]byte, error) {
	buffer := formatter.BufferPool.Get().(*bytes.Buffer)
	defer func() {
		buffer.Reset()

		formatter.BufferPool.Put(buffer)
	}()

	err := formatter.Template.Execute(buffer, map[string]interface{}{
		"datetime": record.Timestamp.Format(time.RFC3339Nano),
		"caller":   record.Caller.String(),
		"level":    record.Level.String(),
		"message":  record.Message,
		"context": TagFunc(func(w io.Writer, tag string) error {
			if len(record.Fields) == 0 {
				if _, err := w.Write([]byte("{}")); err != nil {
					return err
				}

				return nil
			}

			var sb strings.Builder

			sb.WriteString("{")

			hasMore := false
			for _, f := range record.Fields {
				if hasMore {
					sb.WriteString(",")
				}

				switch f.Type {
				case log.FieldBoolType:
					sb.WriteString(fmt.Sprintf("%q:%t", f.Key, f.Value))
				case log.FieldIntType:
					sb.WriteString(fmt.Sprintf("%q:%d", f.Key, f.Value))
				case log.FieldFloat64Type:
					sb.WriteString(fmt.Sprintf("%q:%f", f.Key, f.Value))
				case log.FieldStringType:
					sb.WriteString(fmt.Sprintf("%q:%q", f.Key, f.Value))
				case log.FieldErrorType:
					sb.WriteString(fmt.Sprintf("%q:%q", f.Key, f.Value))
				default:
					sb.WriteString(fmt.Sprintf("%q:%v", f.Key, f.Value))
				}

				hasMore = true
			}

			sb.WriteString("}")

			if _, err := w.Write([]byte(sb.String())); err != nil {
				return err
			}

			return nil
		}),
	})
	if err != nil {
		return nil, err
	}

	buffer.Write([]byte("\n"))

	return buffer.Bytes(), nil
}
