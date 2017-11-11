package traincat

import "gopkg.in/resty.v1"

const URL = "https://api.train.cat"

type (
	Config struct {
		Host  string
		Auth  Auth
		Debug bool
	}

	Auth struct {
		Username        string
		Password        string
		tokenValidUntil int64
	}
)

var config Config

func SetConfig(c Config) {
	config = c

	resty.
		SetHostURL(c.Host).
		SetHeader("Accept", "application/json").
		SetDebug(c.Debug)

	auth()
}
