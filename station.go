package traincat

type (
	Station struct {
		Entity
		Name       string `json:"name"`
		UIC        string `json:"uic"`
		IsRealTime bool   `json:"is_realtime"`
		Hateoas
	}
)

func CGetAllStations() ([]Station, error) {
	c := &Collection{}

	_, err := r(false).
		SetResult(c).
		SetQueryParam(limitPerPage, limitMaxPerPage).
		Get(EndpointStations)

	if err != nil {
		return nil, err
	}

	ss := []Station{}

	for ; err == nil; {
		tmp := []Station{}
		err = c.Embedded.Get(EmbeddedItems, &tmp)

		if err != nil {
			return nil, err
		}

		ss = append(ss, tmp...)
		err = c.Next()
	}

	if err == ErrLastPast {
		err = nil
	}

	return ss, err
}

