package plugins

import "github.com/ibnusurkati/basit/context"

// BasicAuth http basic auth with username and password
type BasicAuth struct {
	Username string
	Password string
}

// Apply http basic auth with username and password
func (b BasicAuth) Apply(ctx *context.Context) {
	ctx.Request.SetBasicAuth(b.Username, b.Password)
}

// Valid http basic auth with username and password valid?
func (b BasicAuth) Valid() bool {
	return b.Username != ""
}
