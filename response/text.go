package response

import "reflect"

func (r *Response) Text(text *interface{}) *Response {
	if r.ctx.HasError() {
		return r
	}

	reflectvalue := reflect.ValueOf(r.buffer.String())

	*text = reflectvalue.Interface()

	return r
}
