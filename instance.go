package basit

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/ibnusurkati/basit/context"
	"github.com/ibnusurkati/basit/plugins"
	"github.com/ibnusurkati/basit/response"
)

type Instance struct {
	Url          string
	Method       string
	ResponseType string
	Query        map[string]string
	Headers      map[string]string
	Data         interface{}
	DataType     string
	Timeout      time.Duration
	TLSTimeout   time.Duration
	DialTimeout  time.Duration
	TLSConfig    *tls.Config
	Transport    *http.Transport
	ProxyURL     string
	ProxyServers map[string]string
	Cookies      []*http.Cookie
	CookiesMap   map[string]string
	BasicAuth    BasicAuth
}

type BasicAuth struct {
	Username string
	Password string
}

func (i *Instance) InitContext() *context.Context {
	ctx := context.New()

	pluginList := []plugins.Setup{
		plugins.URL{Data: i.Url},
		plugins.Method{Data: i.Method},
		plugins.Query{Data: i.Query},
		plugins.Header{Data: i.Headers},
		plugins.Body{Data: i.Data, Type: i.DataType},
		plugins.TLSConfig{Config: i.TLSConfig},
		plugins.Cookies{Data: i.Cookies, Map: i.CookiesMap},
		plugins.BasicAuth{Username: i.BasicAuth.Username, Password: i.BasicAuth.Password},
		plugins.Timeouts{Request: i.Timeout, TLS: i.TLSTimeout, Dial: i.DialTimeout},
		plugins.Proxy{Servers: i.ProxyServers, URL: i.ProxyURL},
		plugins.Transport{RoundTripper: i.Transport},
	}

	for _, plugin := range pluginList {
		if plugin.Valid() {
			plugin.Apply(ctx)
		}
	}

	return ctx
}

func (i *Instance) ExecDo() response.Response {
	ctx := i.InitContext()
	resp := response.New(ctx).Do()
	return *resp
}

func (i *Instance) Exec(data *interface{}) *response.Response {
	ctx := i.InitContext()
	resp := response.New(ctx).Do()

	switch i.ResponseType {
	case "json":
		resp.Json(&data)
	case "text":
		resp.Text(data)
	}

	return resp
}
