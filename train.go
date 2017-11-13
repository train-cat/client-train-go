package traincat

import (
	"errors"
	"net/http"
)

type (
	// Train output of the API
	Train struct {
		Entity
		Code    string `json:"code"`
		Mission string `json:"mission"`
		Hateoas
	}

	// TrainInput for the API
	TrainInput struct {
		Code       *string `json:"code"`
		Mission    *string `json:"mission"`
		TerminusID *uint   `json:"terminus_id"`
	}
)

// TrainExist return true if code is already in database
func TrainExist(code string) (bool, error) {
	resp, err := r(true).
		Head(BuildURI(EndpointTrain, code))

	if err != nil {
		return false, err
	}

	return resp.StatusCode() == http.StatusNoContent, nil
}

// PostTrain add new train to the API
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
