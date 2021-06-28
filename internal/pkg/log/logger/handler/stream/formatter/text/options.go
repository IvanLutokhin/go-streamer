package text

const defaultTemplate = "%datetime% %caller% %level% %message% %context%"

type FormatterOptions struct {
	Template       string `json:"template"`
	SkipUnknownTag bool   `json:"skip_unknown_tag"`
}

func NewDefaultFormatterOptions() *FormatterOptions {
	return &FormatterOptions{
		Template:       defaultTemplate,
		SkipUnknownTag: true,
	}
}
