package plugins

import (
	"net/url"

	"github.com/ibnusurkati/basit/context"
)

// URL http url (host+path+params)
type URL struct {
	Data string
}

// Apply http url (host+path+params)
func (_u URL) Apply(ctx *context.Context) {
	if _u.Data == "" {
		return
	}
	u, err := url.Parse(_u.Data)
	if err != nil {
		return
	}
	ctx.Request.URL = u
}

// Valid http url path like: /api/v1/xx valid?
func (_u URL) Valid() bool {
	return _u.Data != ""
}
