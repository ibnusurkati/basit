package response

import (
	"io"
	"os"
)

func (r *Response) SaveToFile(fileName string) *Response {
	if r.ctx.HasError() {
		return r
	}

	fd, err := os.Create(fileName)
	if err != nil {
		r.ctx.SetError(err)
		goto OUT
	}

	defer r.Close()
	defer fd.Close()

	if _, err = io.Copy(fd, r.buffer); err != nil && err != io.EOF {
		r.ctx.SetError(err)
		goto OUT
	}

OUT:
	return r
}
