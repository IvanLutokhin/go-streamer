package handler

import "encoding/json"

type Options map[string]interface{}

func (options Options) Unmarshal(v interface{}) error {
	if options == nil {
		return nil
	}

	data, err := json.Marshal(options)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}
