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
		Get(EndpointStations)

	if err != nil {
		return nil, err
	}

	var ss []Station

	for err == nil {
		var tmp []Station
		err = c.Embedded.Get(EmbeddedItems, &tmp)

		if err != nil {
			return nil, err
		}

		ss = append(ss, tmp...)

		if c.IsLastPage() {
			break
		}

		err = c.Next()
	}

	return ss, err
}
