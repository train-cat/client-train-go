package traincat

import (
	"errors"
	"encoding/json"
)

type (
	Hateoas struct {
		Embedded Embedded        `json:"_embedded"`
		Links    map[string]Link `json:"_links"`
	}

	Embedded map[string]json.RawMessage

	Link struct {
		Href string `json:"href"`
	}
)

func (e *Embedded) Get(key string, s interface{}) error {
	embedded := map[string]json.RawMessage(*e)

	v, ok := embedded[key]

	if !ok {
		return errors.New("key not found")
	}

	return json.Unmarshal(v, s)
}
