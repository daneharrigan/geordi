package command

import (
	"testing"
	"bytes"
)

func TestParse(t *testing.T) {
	reader := bytes.NewReader([]byte(`SET "key" "value"`))
	name, args := Parse(reader)

	if name != "SET" {
		fatalf(t, "SET", name)
	} 

	if args[0] != "key" {
		fatalf(t, "key", args[0])
	} 

	if args[1] != "value" {
		fatalf(t, "value", args[1])
	} 
}

func fatalf(t *testing.T, a, b string) {
	t.Fatalf("expected %q; got %q", a, b)
}
