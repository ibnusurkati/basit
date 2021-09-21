package plugins

import "github.com/ibnusurkati/basit/context"

type Query struct {
	Data map[string]string
}

func (q Query) Apply(ctx *context.Context) {
	if q.Data == nil {
		return
	}
	query := ctx.Request.URL.Query()
	for k, v := range q.Data {
		query.Set(k, v)
	}
	ctx.Request.URL.RawQuery = query.Encode()
}

func (q Query) Valid() bool {
	return q.Data != nil
}
