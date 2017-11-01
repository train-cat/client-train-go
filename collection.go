package traincat

import "errors"

const (
	limitPerPage    = "_limit_per_page"
	limitMaxPerPage = "100"

	EmbeddedItems = "items"
	LinkNext      = "next"
)

type Collection struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Pages int `json:"pages"`
	Total int `json:"total"`
	Hateoas
}

var ErrLastPast = errors.New("no next page")

func (c *Collection) Next() error {
	if c.Page >= c.Pages {
		return ErrLastPast
	}

	_, err := r(false).
		SetResult(c).
		Get(c.Links[LinkNext].Href)

	return err
}
