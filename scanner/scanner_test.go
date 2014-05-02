package scanner_test

import (
	"github.com/daneharrigan/geordi/scanner"
	"testing"
	"bytes"
)

func TestScan(t *testing.T) {
	b := []byte(`CMD "key" "value" 123 12.3 -1`)
	s := scanner.New(b)
	values := [][]byte{
		[]byte("CMD"),
		[]byte("key"),
		[]byte("value"),
		[]byte("123"),
		[]byte("12.3"),
		[]byte("-1"),
	}

	var i int
	for s.Scan() {
		if !bytes.Equal(s.Bytes(), values[i]) {
			t.Fatalf("expected %q; was %q", s.Bytes(), values[i])
		}
		i++
	}

	if i < 6 {
		t.Fatalf("expected %d; was %d", 6, i)
	}

	if s.Err() != nil {
		t.Fatalf("expected %v; was %q", nil, s.Err())
	}
}
