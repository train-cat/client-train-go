package traincat

import (
	"encoding/json"
	"errors"
)

type (
	// Hateoas structure sent by the API
	Hateoas struct {
		Embedded Embedded        `json:"_embedded"`
		Links    map[string]Link `json:"_links"`
	}

	// Embedded alias to map[string]json.RawMessage
	Embedded map[string]json.RawMessage

	// Link structure sent by the API
	Link struct {
		Href string `json:"href"`
	}
)

// Get one value from _embedded key
func (e *Embedded) Get(key string, s interface{}) error {
	embedded := map[string]json.RawMessage(*e)

	v, ok := embedded[key]

	if !ok {
		return errors.New("key not found")
	}

	return json.Unmarshal(v, s)
}
