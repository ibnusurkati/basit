package context

import (
	"net/http"
)

type Context struct {
	Client   *http.Client   `json:"client,omitempty"`
	Request  *http.Request  `json:"request,omitempty"`
	Response *http.Response `json:"response,omitempty"`

	TimeTrace Time
	err       error
}

func (c *Context) GetRequest() *http.Request {
	return c.Request
}

func New() *Context {
	request, _ := http.NewRequest("", "", nil)

	request.Header.Set("User-Agent", "basit")
	return &Context{
		Request: request,
		Client: &http.Client{
			Transport: http.DefaultTransport,
		},
	}
}
