package traincat

import (
	"errors"
	"net/http"
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
		Head(BuildURI(EndpointStationTrainStop, stationID, code))

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
		Post(BuildURI(EndpointStationTrainStop, stationID, code))

	if resp.StatusCode() != http.StatusCreated {
		return nil, errors.New(string(resp.Body()))
	}

	return s, err
}
