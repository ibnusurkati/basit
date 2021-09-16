package plugins

import (
	"net"
	"net/http"
	"time"

	"github.com/ibnusurkati/basit/context"
)

var (
	DialTimeout = 30 * time.Second

	DialKeepAlive = 30 * time.Second

	TLSHandshakeTimeout = 10 * time.Second

	RequestTimeout = 60 * time.Second

	DefaultDialer = &net.Dialer{
		Timeout:   DialTimeout,
		KeepAlive: DialKeepAlive,
	}
)

type Timeouts struct {
	Request time.Duration

	TLS time.Duration

	Dial time.Duration

	KeepAlive time.Duration
}

func (to Timeouts) Apply(ctx *context.Context) {
	if to.Request == 0 {
		to.Request = RequestTimeout
	}
	ctx.Client.Timeout = to.Request

	transport, ok := ctx.Client.Transport.(*http.Transport)
	if !ok {
		return
	}

	if to.TLS == 0 {
		to.TLS = TLSHandshakeTimeout
	}
	transport.TLSHandshakeTimeout = to.TLS

	if to.Dial == 0 {
		to.Dial = DialTimeout
	}
	if to.KeepAlive == 0 {
		to.KeepAlive = DialKeepAlive
	}

	transport.Dial = (&net.Dialer{
		Timeout:   to.Dial,
		KeepAlive: to.KeepAlive,
	}).Dial

	ctx.Client.Transport = transport
}
func (to Timeouts) Valid() bool {
	if to.Request == 0 && to.TLS == 0 &&
		to.Dial == 0 && to.KeepAlive == 0 {
		return false
	}
	return true
}
