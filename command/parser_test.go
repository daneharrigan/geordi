package command_test

import (
	"testing"
	"bytes"
)

func TestParser(t *testing.T) {
	cmd1 := []byte(`SET "name" "test value"`)
	//cmd2 := []byte(`SET "name" "te\"s\"t"`)
	//cmd3 := []byte(`SET "age" 28`)
	//cmd4 := []byte("\n\n    "+`GET "age"`) // ignore white space

	var name string
	var args [][]byte

	name, args = parser(cmd1)
	if name != "SET" {
		fatalf(t, "SET", name)
	}

	if !bytes.Equal([]byte("name"), args[0]) {
		fatalf(t, "name", args[0])
	}
}

func fatalf(t *testing.T, data ...interface{}) {
	t.Fatalf("expected %q; got %q", data...)
}

// parser.go

const (
	CR = '\r'
	LF = '\n'
	Esc = '\\'
)

func parser(b []byte) (name string, args [][]byte) {
	var quoted, escaped bool
	var i, s, f int
	var c byte
garbage:
	c = b[i]
	switch {
	case c == CR && (quoted || escaped):
		fallthrough
	case c == LF && (quoted || escaped):
		i++
		goto content
	}
content:
	c = b[i]
	switch {
	case c == Esc:
	}
finish:
	return
}
