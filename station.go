package traincat

import (
	"fmt"

	"github.com/train-cat/client-train-go/filters"
)

type (
	// Station output of the API
	Station struct {
		Entity
		Name       string `json:"name"`
		UIC        string `json:"uic"`
		IsRealTime bool   `json:"is_realtime"`
		Hateoas
	}
)

// GetStation return one station
func GetStation(stationID int) (*Station, error) {
	s := &Station{}

	_, err := r(false).SetResult(s).
		Get(fmt.Sprintf(EndpointStation, stationID))

	return s, err
}

// CGetAllStations get all pages and return all stations
func CGetAllStations(f *filters.Station) ([]Station, error) {
	c := &Collection{}

	req := r(false).
		SetResult(c)

	_, err := filters.Apply(req, f).Get(EndpointStations)

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
