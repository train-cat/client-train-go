package traincat

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/train-cat/client-train-go/filters"
)

type (
	// Stop output of the API
	Stop struct {
		Entity
		Schedule string `json:"schedule"`
		OnWeek   bool   `json:"on_week"`
		Hateoas
	}

	// StopInput for the API
	StopInput struct {
		Schedule string `json:"schedule"`
		IsWeek   bool   `json:"is_week"`
	}
)

// StopExist return true if one stop exist (stationID + code)
func StopExist(stationID uint, code string) (bool, error) {
	resp, err := r(true).
		Head(BuildURI(EndpointStationTrainStops, stationID, code))

	if err != nil {
		return false, err
	}

	return resp.StatusCode() == http.StatusNoContent, nil
}

// PostStop add new stop to the API
func PostStop(stationID uint, code string, i StopInput) (*Stop, error) {
	refreshAuth()

	s := &Stop{}

	resp, err := r(true).
		SetBody(i).
		SetResult(s).
		Post(BuildURI(EndpointStationTrainStops, stationID, code))

	if resp.StatusCode() != http.StatusCreated {
		return nil, errors.New(string(resp.Body()))
	}

	return s, err
}

// CGetStops return all stops available for one station
func CGetStops(stationID uint, f *filters.Stop) ([]Stop, error) {
	c := &Collection{}

	req := r(false).SetResult(c)

	_, err := filters.Apply(req, f).Get(fmt.Sprintf(EndpointStationStop, stationID))

	if err != nil {
		return nil, err
	}

	var stops []Stop

	err = c.Embedded.Get(EmbeddedItems, &stops)

	return stops, err
}
