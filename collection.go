package traincat

import "errors"

// All keys can be found in hateoas collection
const (
	EmbeddedItems = "items"
	LinkFirst     = "first"
	LinkLast      = "last"
	LinkPrevious  = "previous"
	LinkNext      = "next"
)

// Collection represent structure return by the API
type Collection struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Pages int `json:"pages"`
	Total int `json:"total"`
	Hateoas
}

// Error relevant to collection
var (
	ErrFirstPage = errors.New("no previous page")
	ErrLastPage  = errors.New("no next page")
)

// First load the first page
func (c *Collection) First() error {
	_, err := r(false).
		SetResult(c).
		Get(c.Links[LinkFirst].Href)

	return err
}

// Last load the last page
func (c *Collection) Last() error {
	_, err := r(false).
		SetResult(c).
		Get(c.Links[LinkLast].Href)

	return err
}

// Previous load previous page
func (c *Collection) Previous() error {
	if c.IsFirstPage() {
		return ErrFirstPage
	}

	_, err := r(false).
		SetResult(c).
		Get(c.Links[LinkPrevious].Href)

	return err
}

// Next load next page
func (c *Collection) Next() error {
	if c.IsLastPage() {
		return ErrLastPage
	}

	_, err := r(false).
		SetResult(c).
		Get(c.Links[LinkNext].Href)

	return err
}

// IsFirstPage return true if collection is currently on the first page
func (c *Collection) IsFirstPage() bool {
	return c.Page == 1
}

// IsLastPage return true if collection is currently on the last page
func (c *Collection) IsLastPage() bool {
	return c.Page >= c.Pages
}
