package filters

type Pagination struct {
	Page       *int `query:"_page"`
	MaxPerPage *int `query:"_max_per_page"`
}
