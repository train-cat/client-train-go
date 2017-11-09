package traincat

import "strconv"

type (
	Alert struct {
		Entity
		ActionID uint `json:"action_id"`
		Hateoas
	}

	AlertFilter struct {
		CodeTrain string
		StationID int
	}
)

func CGetAllAlerts(f AlertFilter) ([]Alert, error) {
	c := &Collection{}

	req := r(false).
		SetResult(c).
		SetQueryParam(limitPerPage, limitMaxPerPage)

	if f.CodeTrain != "" {
		req.SetQueryParam("code_train", f.CodeTrain)
	}

	if f.StationID != 0 {
		req.SetQueryParam("station_id", strconv.Itoa(f.StationID))
	}

	_, err := req.Get(EndpointAlert)

	if err != nil {
		return nil, err
	}

	as := []Alert{}

	for err == nil {
		tmp := []Alert{}
		err = c.Embedded.Get(EmbeddedItems, &tmp)

		if err != nil {
			return nil, err
		}

		as = append(as, tmp...)
		err = c.Next()
	}

	if err == ErrLastPast {
		err = nil
	}

	return as, err
}
