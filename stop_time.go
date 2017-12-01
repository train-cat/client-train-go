package traincat

import (
	"fmt"

	"github.com/train-cat/client-train-go/filters"
)

type (
	// StopTime output of the API
	StopTime struct {
		Entity
		Schedule string `json:"schedule"`
		Hateoas
	}

	// StopInput for the API
	StopInput struct {
		Schedule string `json:"schedule"`
		IsWeek   bool   `json:"is_week"`
	}
)

// CGetStops return all stopsTime available for one station
func CGetStopsTime(stationID uint, f *filters.StopTime) ([]StopTime, error) {
	c := &Collection{}

	req := r(false).SetResult(c)

	_, err := filters.Apply(req, f).Get(fmt.Sprintf(EndpointStationStopsTime, stationID))

	if err != nil {
		return nil, err
	}

	var stops []StopTime

	err = c.Embedded.Get(EmbeddedItems, &stops)

	return stops, err
}
