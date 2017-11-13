package traincat

// Type of action available
const (
	ActionTypeMessenger = "messenger"
)

// Action output from the API
type Action struct {
	Entity
	Type string            `json:"type"`
	Data map[string]string `json:"data"`
	Hateoas
}
