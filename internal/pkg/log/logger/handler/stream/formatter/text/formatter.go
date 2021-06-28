package text

import (
	"bytes"
	"fmt"
	"github.com/IvanLutokhin/go-streamer/internal/pkg/log"
	"github.com/IvanLutokhin/go-streamer/internal/pkg/log/logger/handler"
	"github.com/IvanLutokhin/go-streamer/internal/pkg/log/logger/handler/stream/formatter"
	"io"
	"strings"
	"sync"
	"time"
)

type Formatter struct {
	template   *Template
	bufferPool *sync.Pool
}

func NewFormatter(options handler.Options) (formatter formatter.Formatter, err error) {
	var formatterOptions *FormatterOptions
	if formatterOptions, err = ParseFormatterOptions(options); err != nil {
		return
	}

	var t *Template
	if t, err = ParseTemplate(formatterOptions.Template, formatterOptions.SkipUnknownTag); err != nil {
		return
	}

	formatter = &Formatter{
		template: t,
		bufferPool: &sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
	}

	return
}

func ValidateFormatter(options handler.Options) (err error) {
	var formatterOptions *FormatterOptions
	if formatterOptions, err = ParseFormatterOptions(options); err != nil {
		return
	}

	if _, err = ParseTemplate(formatterOptions.Template, formatterOptions.SkipUnknownTag); err != nil {
		return
	}

	return
}

func ParseFormatterOptions(options handler.Options) (*FormatterOptions, error) {
	formatterOptions := NewDefaultFormatterOptions()

	if err := options.Unmarshal(formatterOptions); err != nil {
		return nil, err
	}

	return formatterOptions, nil
}

func (formatter *Formatter) Format(record handler.Record) ([]byte, error) {
	buffer := formatter.bufferPool.Get().(*bytes.Buffer)
	defer func() {
		buffer.Reset()

		formatter.bufferPool.Put(buffer)
	}()

	err := formatter.template.Execute(buffer, map[string]interface{}{
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

func init() {
	formatter.RegisterFormatter("text", NewFormatter, ValidateFormatter)
}
