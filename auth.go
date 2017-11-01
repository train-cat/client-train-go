package traincat

import (
	"errors"
	"net/http"
	"time"

	"gopkg.in/resty.v1"
)

type (
	credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	token struct {
		Token string `json:"token"`
	}
)

func auth() error {
	a := credentials{config.Username, config.Password}
	t := &token{}

	resp, err := r(true).
		SetBody(a).
		SetResult(t).
		Post(EndpointLogin)

	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		return errors.New("bad credentials")
	}

	resty.SetAuthToken(t.Token)

	config.tokenValidUntil = time.Now().Add(time.Hour * 71).Unix()

	return nil
}

func refreshAuth() error {
	if time.Now().Unix() > config.tokenValidUntil {
		return auth()
	}

	return nil
}
