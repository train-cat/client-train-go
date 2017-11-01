package traincat

import (
	"errors"
	"fmt"
	"net/http"
)

type (
	Stop struct {
		Entity
		Schedule string `json:"schedule"`
		OnWeek   bool   `json:"on_week"`
		Hateoas
	}

	StopInput struct {
		Schedule string `json:"schedule"`
		IsWeek   bool   `json:"is_week"`
	}
)

func StopExist(stationID uint, code string) (bool, error) {
	resp, err := r(true).
		Head(fmt.Sprintf(EndpointStationTrainStop, stationID, code))

	if err != nil {
		return false, err
	}

	return resp.StatusCode() == http.StatusNoContent, nil
}

func PostStop(stationID uint, code string, i StopInput) (*Stop, error) {
	refreshAuth()

	s := &Stop{}

	resp, err := r(true).
		SetBody(i).
		SetResult(s).
		Post(fmt.Sprintf(EndpointStationTrainStop, stationID, code))

	if resp.StatusCode() != http.StatusCreated {
		return nil, errors.New(string(resp.Body()))
	}

	return s, err
}