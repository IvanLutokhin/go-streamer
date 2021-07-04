package text

import (
	"bytes"
	"encoding/json"
	handlerConfig "github.com/IvanLutokhin/go-streamer/internal/app/streamer/config/logger/handlers/stream"
	"github.com/IvanLutokhin/go-streamer/internal/pkg/logger/handlers/stream"
	"github.com/IvanLutokhin/go-streamer/internal/pkg/logger/handlers/stream/formatters/text"
	"sync"
)

func NewFormatter(options map[string]interface{}) (formatter stream.Formatter, err error) {
	var formatterConfig *FormatterConfig
	if formatterConfig, err = ParseFormatterConfig(options); err != nil {
		return
	}

	var t *text.Template
	if t, err = text.ParseTemplate(formatterConfig.Template, formatterConfig.SkipUnknownTag); err != nil {
		return
	}

	formatter = &text.Formatter{
		Template: t,
		BufferPool: &sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
	}

	return
}

func ValidateFormatter(options map[string]interface{}) (err error) {
	var formatterConfig *FormatterConfig
	if formatterConfig, err = ParseFormatterConfig(options); err != nil {
		return
	}

	if _, err = text.ParseTemplate(formatterConfig.Template, formatterConfig.SkipUnknownTag); err != nil {
		return
	}

	return
}

func ParseFormatterConfig(options map[string]interface{}) (*FormatterConfig, error) {
	formatterConfig := NewDefaultFormatterConfig()

	if options == nil {
		return formatterConfig, nil
	}

	data, err := json.Marshal(options)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, formatterConfig); err != nil {
		return nil, err
	}

	return formatterConfig, nil
}

func init() {
	handlerConfig.RegisterFormatter("text", NewFormatter, ValidateFormatter)
}
