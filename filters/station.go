package filters

// Station filters
type Station struct {
	Pagination
	Name *string `query:"name"`
}
