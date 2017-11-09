package traincat

const (
	ActionTypeMessenger = "messenger"
)

type Action struct {
	Entity
	Type string            `json:"type"`
	Data map[string]string `json:"data"`
	Hateoas
}
