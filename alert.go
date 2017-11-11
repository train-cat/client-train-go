package traincat

import "github.com/train-cat/client-train-go/filters"

type (
	Alert struct {
		Entity
		ActionID uint `json:"action_id"`
		Hateoas
	}
)

func CGetAllAlerts(f *filters.Alert) ([]Alert, error) {
	c := &Collection{}

	req := r(false).
		SetResult(c)

	_, err := filters.Apply(req, f).Get(EndpointAlert)

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
