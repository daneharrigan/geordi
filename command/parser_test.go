package command_test

import (
	"testing"
	"fmt"
	"bytes"
)

func TestParser(t *testing.T) {
	cmd1 := []byte(`SET "name" "test value"`)
	//cmd2 := []byte(`SET "name" "te\"s\"t"`)
	//cmd3 := []byte(`SET "age" 28`)
	//cmd4 := []byte("\n\n    "+`GET "age"`) // ignore white space

	var name string
	var args [][]byte

	println("at=TestParser")
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

func parser(b []byte) (name string, args [][]byte) {
println("at=start")
	var quoted, escaped bool
	var i, s, f int
	var c byte
garbage:
	println("at=garbage")
	c = b[i]
	switch c {
	case '\n', '\r', '\t', ' ':
		i++
		goto garbage
	default:
		goto content
	}
content:
	fmt.Printf("at=content i=%d\n", i)
	if len(b) == i {
		return
	}

	c = b[i]
	switch c {
	case '"':
		i++
		if !quoted {
			quoted = true
			s = i
		} else if quoted && escaped {
			quoted = true
			escaped = false
		} else if quoted {
			f = i
			quoted = false
		}

		goto content
	case '\\':
		escaped = !escaped
		i++
		goto content
	case ' ':
		if escaped {
			i++
		} else {
			f = i
			if name == "" {
				fmt.Printf("s=%d f=%d\n", s, f)
				name = string(b[s:f])
				return
			} else {
				args = append(args, b[s:f])
			}
		}

		goto content
	default:
		i++
		goto content
	}

	return
}
