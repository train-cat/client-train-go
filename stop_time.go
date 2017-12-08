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

// GetStation return station associate
func (st StopTime) GetStation() (*Station, error) {
	s := &Station{}

	_, err := r(false).SetResult(s).
		Get(st.Links["station"].Href)

	return s, err
}

// GetTerminus return terminus of the trip
func (st StopTime) GetTerminus() (*Station, error) {
	s := &Station{}

	_, err := r(false).SetResult(s).
		Get(fmt.Sprintf("%s%s", st.Links["trip"].Href, endpointTerminus))

	return s, err
}
