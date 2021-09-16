package plugins

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/url"

	"github.com/ibnusurkati/basit/context"
	"gopkg.in/yaml.v2"
)

type Body struct {
	Data interface{}
	Type string
}

// Apply json body
func (b Body) Apply(ctx *context.Context) {
	switch b.Type {
	case "text":
		b.Text(ctx)
	case "json":
		b.Json(ctx)
	case "xml":
		b.Xml(ctx)
	case "yaml":
		b.Xml(ctx)
	case "urlencode":
		b.UrlEncode(ctx)
	}
}

func (b Body) Text(ctx *context.Context) {
	bBytes := bytes.NewReader([]byte(b.Data.(string)))
	rc, ok := io.Reader(bBytes).(io.ReadCloser)
	if !ok && bBytes != nil {
		rc = io.NopCloser(bBytes)
	}

	ctx.Request.Body = rc
	ctx.Request.ContentLength = int64(bytes.NewBufferString(b.Data.(string)).Len())
}

func (b Body) Json(ctx *context.Context) {
	buf := &bytes.Buffer{}

	switch b.Data.(type) {
	case string:
		buf.WriteString(b.Data.(string))
	case []byte:
		buf.Write(b.Data.([]byte))
	default:
		if err := json.NewEncoder(buf).Encode(b.Data); err != nil {
			ctx.SetError(fmt.Errorf("unknown json encoded type: %T", b.Data))
			return
		}
	}

	ctx.Request.Body = io.NopCloser(buf)
	ctx.Request.ContentLength = int64(buf.Len())
}

func (b Body) Xml(ctx *context.Context) {
	buf := &bytes.Buffer{}

	switch b.Data.(type) {
	case string:
		buf.WriteString(b.Data.(string))
	case []byte:
		buf.Write(b.Data.([]byte))
	default:
		if err := xml.NewEncoder(buf).Encode(b.Data); err != nil {
			ctx.SetError(fmt.Errorf("unknown xml encoded type: %T", b.Data))
			return
		}
	}

	ctx.Request.Body = io.NopCloser(buf)
	ctx.Request.ContentLength = int64(buf.Len())
}

func (b Body) Yaml(ctx *context.Context) {
	buf := &bytes.Buffer{}

	switch b.Data.(type) {
	case string:
		buf.WriteString(b.Data.(string))
	case []byte:
		buf.Write(b.Data.([]byte))
	default:
		if err := yaml.NewEncoder(buf).Encode(b.Data); err != nil {
			ctx.SetError(fmt.Errorf("unknown yaml encoded type: %T", b.Data))
			return
		}
	}

	ctx.Request.Body = io.NopCloser(buf)
	ctx.Request.ContentLength = int64(buf.Len())
}

func (b Body) UrlEncode(ctx *context.Context) {
	buf := &bytes.Buffer{}

	switch b.Data.(type) {
	case string:
		buf.WriteString(b.Data.(string))
	case []byte:
		buf.Write(b.Data.([]byte))
	case map[string]string:
		data := make(url.Values)
		for k, v := range b.Data.(map[string]string) {
			data.Set(k, v)
		}
		buf.WriteString(data.Encode())
	case map[string][]string:
		buf.WriteString(url.Values(b.Data.(map[string][]string)).Encode())
	case url.Values:
		buf.WriteString(b.Data.(url.Values).Encode())
	default:
		ctx.SetError(fmt.Errorf("unknown urlencoded type: %T", b.Data))
		return
	}

	ctx.Request.Body = io.NopCloser(buf)
	ctx.Request.ContentLength = int64(buf.Len())
}

// Valid json body valid?
func (b Body) Valid() bool {
	data := b.Type
	dataTypes := []string{
		"text",
		"json",
		"xml",
		"yaml",
		"urlencode",
	}
	for _, dataType := range dataTypes {
		if data == dataType {
			return true
		}
	}
	return false
}
