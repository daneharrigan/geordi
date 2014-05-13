package responder

import (
	"bytes"
	"io"
	"fmt"
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

func (r *Responder) SetError() {
	r.buffer.WriteByte('-')
}

func (r *Responder) SetSuccess() {
	r.buffer.WriteByte('+')
}

func (r *Responder) Write(b []byte) {
	r.buffer.WriteByte(' ')
	r.buffer.Write(b)
}

func (r *Responder) WriteString(s string) {
	r.buffer.WriteByte(' ')
	r.buffer.WriteByte('"')
	r.buffer.WriteString(s)
	r.buffer.WriteByte('"')
}

func (r *Responder) WriteError(err error) {
	r.SetError()
	r.WriteString(err.Error())
}

func (r *Responder) WriteInt(n int64) {
	fmt.Fprintf(r.buffer, " %d", n)
}

func (r *Responder) WriteFloat(f float64) {
	fmt.Fprintf(r.buffer, " %f", f)
}

func (r *Responder) Flush() error {
	r.buffer.WriteByte('\r')
	_, err := r.buffer.WriteTo(r.writer)
	r.buffer.Reset()
	return err
}
