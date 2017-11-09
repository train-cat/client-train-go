package traincat

import (
	"fmt"

	"gopkg.in/resty.v1"
)

func r(contentType bool) *resty.Request {
	r := resty.R()

	if contentType {
		return r.SetHeader("Content-Type", "application/json")
	}

	return r
}

// BuildURI return URI formatted
func BuildURI(endpoint string, a ...interface{}) string {
	return fmt.Sprintf(endpoint, a...)
}
