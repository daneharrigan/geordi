package responder

import (
	"bytes"
	"io"
)

type Responder struct {
	bytes  []byte
	writer io.Writer
	buffer *bytes.Buffer
}

func New(writer io.Writer) *Responder {
	responder := &Responder{writer: writer}
	responder.buffer = bytes.NewBuffer(responder.bytes)
	return responder
}

func (r *Responder) WriteError() {
	r.buffer.WriteByte('-')
}

func (r *Responder) WriteSuccess() {
	r.buffer.WriteByte('+')
}

func (r *Responder) Write(b []byte) {
	r.buffer.Write(b)
}

func (r *Responder) WriteString(s string) {
	r.buffer.WriteString(s)
}

func (r *Responder) Flush() error {
	r.buffer.WriteByte('\r')
	_, err := r.buffer.WriteTo(r.writer)
	r.buffer.Reset()
	return err
}
