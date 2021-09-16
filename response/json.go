package response

import (
	"encoding/json"
	"io"
)

func (r *Response) Json(userStruct interface{}) *Response {
	if r.ctx.HasError() {
		return r
	}

	jsonDecoder := json.NewDecoder(r.buffer)

	defer r.Close()

	if err := jsonDecoder.Decode(userStruct); err != nil && err != io.EOF {
		r.ctx.SetError(err)
	}

	return r
}
