package traincat

import "gopkg.in/resty.v1"

func r(contentType bool) *resty.Request {
	r := resty.R()

	if contentType {
		return r.SetHeader("Content-Type", "application/json")
	}

	return r
}
