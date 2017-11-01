package traincat

import "gopkg.in/resty.v1"

type Config struct {
	Host            string
	Username        string
	Password        string
	tokenValidUntil int64
}

var config Config

func SetConfig(c Config) {
	config = c

	resty.
		SetHostURL(c.Host).
		SetHeader("Accept", "application/json")

	auth()
}
