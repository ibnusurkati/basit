package plugins

import (
	"github.com/ibnusurkati/basit/context"
)

// Header http header
type Header struct {
	Data map[string]string
}

// Apply apply http headers
func (h Header) Apply(ctx *context.Context) {
	for k, v := range h.Data {
		ctx.Request.Header.Set(k, v)
	}
}

// Valid user agent in header valid?
func (h Header) Valid() bool {
	return h.Data != nil
}
