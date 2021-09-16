package plugins

import (
	"github.com/ibnusurkati/basit/context"
)

// Method http method: GET, POST, DELETE ...
type Method struct {
	Data string
}

// Apply http method: GET, POST, DELETE ...
func (m Method) Apply(ctx *context.Context) {
	ctx.Request.Method = m.Data
}

// Valid  http method: GET, POST, DELETE ... valid?
func (m Method) Valid() bool {
	data := m.Data
	methods := []string{
		"OPTIONS",
		"GET",
		"HEAD",
		"POST",
		"PUT",
		"DELETE",
		"TRACE",
		"CONNECT",
		"PATCH",
	}
	for _, method := range methods {
		if data == method {
			return true
		}
	}
	return false
}
