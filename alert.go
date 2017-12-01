package traincat

import (
	"errors"
	"net/http"

	"github.com/train-cat/client-train-go/filters"
)

type (
	// Alert output from the API
	Alert struct {
		Entity
		ActionID uint `json:"action_id"`
		Hateoas
	}

	// AlertInput for the API
	AlertInput struct {
		ActionID uint `json:"action_id"`
	}
)

// CGetAllAlerts get all pages and return all alerts
func CGetAllAlerts(f *filters.Alert) ([]Alert, error) {
	c := &Collection{}

	req := r(false).
		SetResult(c)

	_, err := filters.Apply(req, f).Get(EndpointAlerts)

	if err != nil {
		return nil, err
	}

	var as []Alert

	for err == nil {
		var tmp []Alert
		err = c.Embedded.Get(EmbeddedItems, &tmp)

		if err != nil {
			return nil, err
		}

		as = append(as, tmp...)

		if c.IsLastPage() {
			break
		}

		err = c.Next()
	}

	return as, err
}

// PostAlert add new alert to the API
func PostAlert(stationID int, stopTimeID int, i AlertInput) (*Alert, error) {
	refreshAuth()

	a := &Alert{}

	resp, err := r(true).
		SetBody(i).
		SetResult(a).
		Post(BuildURI(EndpointStationAlerts, stationID, stopTimeID))

	if resp.StatusCode() != http.StatusCreated {
		return nil, errors.New(string(resp.Body()))
	}

	return a, err
}
