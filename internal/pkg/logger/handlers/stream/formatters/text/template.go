package text

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
)

func getStartTag() []byte {
	return []byte("%")
}

func getEndTag() []byte {
	return []byte("%")
}

type TagFunc func(w io.Writer, tag string) error

type Template struct {
	template       string
	skipUnknownTag bool
	texts          [][]byte
	tags           []string
}

func ParseTemplate(template string, skipUnknownTag bool) (*Template, error) {
	if strings.TrimSpace(template) == "" {
		return nil, errors.New("template: could not be empty")
	}

	s := []byte(template)
	a := getStartTag()
	b := getEndTag()

	tagsCount := bytes.Count(s, a)
	if tagsCount == 0 {
		return nil, errors.New("template: tags not defined")
	}

	t := &Template{
		template:       template,
		skipUnknownTag: skipUnknownTag,
		texts:          make([][]byte, 0, tagsCount+1),
		tags:           make([]string, 0, tagsCount),
	}

	for {
		n := bytes.Index(s, a)
		if n < 0 {
			t.texts = append(t.texts, s)

			break
		}
		t.texts = append(t.texts, s[:n])
		s = s[n+len(a):]

		n = bytes.Index(s, b)
		if n < 0 {
			return nil, errors.New("template: cannot find end tag")
		}
		t.tags = append(t.tags, string(s[:n]))
		s = s[n+len(b):]
	}

	return t, nil
}

func (t *Template) ExecuteFunc(w io.Writer, f TagFunc) error {
	n := len(t.texts) - 1
	for i := 0; i < n; i++ {
		if _, err := w.Write(t.texts[i]); err != nil {
			return err
		}

		if err := f(w, t.tags[i]); err != nil {
			return err
		}
	}

	if _, err := w.Write(t.texts[n]); err != nil {
		return err
	}

	return nil
}

func (t *Template) Execute(w io.Writer, m map[string]interface{}) error {
	return t.ExecuteFunc(w, func(w io.Writer, tag string) error {
		return stdTagFunc(w, tag, m, t.skipUnknownTag)
	})
}

func stdTagFunc(w io.Writer, tag string, m map[string]interface{}, skipUnknownTag bool) (err error) {
	v, ok := m[tag]
	if !ok && !skipUnknownTag {
		if _, err = w.Write(getStartTag()); err != nil {
			return
		}

		if _, err = w.Write([]byte(tag)); err != nil {
			return
		}

		if _, err = w.Write(getEndTag()); err != nil {
			return
		}

		return
	}

	if v == nil {
		return
	}

	switch value := v.(type) {
	case []byte:
		_, err = w.Write(value)
	case string:
		_, err = w.Write([]byte(value))
	case TagFunc:
		err = value(w, tag)
	default:
		err = fmt.Errorf("template: tag %q contains unexpected value type=%#v. Expected []byte, string or TagFunc", tag, v)
	}

	return
}
