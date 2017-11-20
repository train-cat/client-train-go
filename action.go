package traincat

import (
	"errors"
	"net/http"
)

// Type of action available
const (
	ActionTypeTelegram  = "telegram"
	ActionTypeMessenger = "messenger"
)

type (
	// Action output from the API
	Action struct {
		Entity
		Type string            `json:"type"`
		Data map[string]string `json:"data"`
		Hateoas
	}

	// ActionInput for the API
	ActionInput struct {
		Type string            `json:"type"`
		Data map[string]string `json:"data"`
	}
)

// PostAction add nex action to the API
func PostAction(i ActionInput) (*Action, error) {
	refreshAuth()

	a := &Action{}

	resp, err := r(true).
		SetBody(i).
		SetResult(a).
		Post(EndpointActions)

	if resp.StatusCode() != http.StatusCreated {
		return nil, errors.New(string(resp.Body()))
	}

	return a, err
}
