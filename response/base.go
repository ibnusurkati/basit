package response

import (
	"bytes"
	"fmt"
	"io"

	"github.com/ibnusurkati/basit/context"
)

type Response struct {
	ctx    *context.Context
	buffer *bytes.Buffer
	done   bool
}

func New(ctx *context.Context) *Response {
	return &Response{ctx: ctx, buffer: bytes.NewBuffer([]byte{})}
}

func (r *Response) OK() bool {
	return !r.ctx.HasError()
}

func (r *Response) Error() error {
	return r.ctx.Error()
}

func (r *Response) Close() *Response {
	if r.ctx.HasError() {
		return r
	}
	if _, err := io.Copy(io.Discard, r.ctx.Response.Body); err != nil {
		r.ctx.SetError(err)
		return r
	}
	r.ctx.SetError(r.ctx.Response.Body.Close())
	return r
}

func (r *Response) ContentType() string {
	return r.ctx.Response.Header.Get("Content-Type")
}

func (r *Response) Code() int {
	return r.ctx.Response.StatusCode
}

func (r *Response) Bytes() []byte {
	if r.buffer.Len() == 0 {
		return nil
	}
	return r.buffer.Bytes()
}

func (r *Response) String() string {
	return r.buffer.String()
}

func (r *Response) TimeTrace() context.Time {
	return r.ctx.TimeTrace
}

func (r *Response) Do() *Response {
	if r.done || r.ctx.HasError() {
		goto OUT
	}

	if err := r.ctx.TraceDo(); err != nil {
		r.ctx.SetError(err)
		fmt.Println(err)
		goto OUT
	}

	if r.buffer.Len() != 0 {
		goto OUT

	}

	if r.ctx.Response.ContentLength == 0 {
		goto OUT
	}

	if r.ctx.Response.ContentLength > 0 {
		r.buffer.Grow(int(r.ctx.Response.ContentLength))
	}

	if _, err := io.Copy(r.buffer, r.ctx.Response.Body); err != nil && err != io.EOF {
		r.ctx.SetError(err)
		r.ctx.Response.Body.Close()
	}
	r.done = true
OUT:
	return r
}
