package traincat

import (
	"errors"
	"fmt"
	"net/http"
)

type (
	Train struct {
		Entity
		Code    string `json:"code"`
		Mission string `json:"mission"`
		Hateoas
	}

	TrainInput struct {
		Code       *string `json:"code"`
		Mission    *string `json:"mission"`
		TerminusID *uint   `json:"terminus_id"`
	}
)

func TrainExist(code string) (bool, error) {
	resp, err := r(true).
		Head(fmt.Sprintf(EndpointTrain, code))

	if err != nil {
		return false, err
	}

	return resp.StatusCode() == http.StatusNoContent, nil
}

func PostTrain(i TrainInput) (*Train, error) {
	refreshAuth()

	t := &Train{}

	resp, err := r(true).
		SetBody(i).
		SetResult(t).
		Post(EndpointTrains)

	if resp.StatusCode() != http.StatusCreated {
		return nil, errors.New(string(resp.Body()))
	}

	return t, err
}
