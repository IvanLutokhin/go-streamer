package text

type FormatterConfig struct {
	Template       string `json:"template"`
	SkipUnknownTag bool   `json:"skip_unknown_tag"`
}

func NewDefaultFormatterConfig() *FormatterConfig {
	return &FormatterConfig{
		Template:       "%datetime% %caller% %level% %message% %context%",
		SkipUnknownTag: true,
	}
}
