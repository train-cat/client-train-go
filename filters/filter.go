package filters

import (
	"reflect"

	"gopkg.in/resty.v1"
	"strconv"
)

const query = "query"

// Apply filter on request
func Apply(r *resty.Request, f interface{}) *resty.Request {
	v := reflect.Indirect(reflect.ValueOf(f))

	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag.Get(query)
		f := v.Field(i)

		if isValidTag(tag) && !f.IsNil() {
			r.SetQueryParam(tag, castToString(f))
		}
	}

	return r
}

func castToString(f reflect.Value) string {
	switch f.Kind() {
	case reflect.Int:
		return strconv.Itoa(int(f.Int()))
	case reflect.String:
		return f.String()
	case reflect.Ptr:
		return castToString(f.Elem())
	}

	return ""
}

func isValidTag(tag string) bool {
	return tag != "-" && tag != ""
}
