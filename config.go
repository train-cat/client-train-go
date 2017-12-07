package traincat

import "gopkg.in/resty.v1"

// URL of the API
const URL = "https://api.train.cat"

type (
	// Config used by the client for call the API
	Config struct {
		Host  string
		Auth  Auth
		Debug bool
	}

	// Auth structure for log the user
	Auth struct {
		Username        string
		Password        string
		tokenValidUntil int64
	}
)

var config Config

// SetConfig apply the config to the current client
func SetConfig(c Config) {
	config = c

	resty.
		SetHostURL(c.Host).
		SetHeader("Accept", "application/json").
		SetDebug(c.Debug)

	if c.Auth.Username != "" && c.Auth.Password != "" {
		auth()
	}
}
