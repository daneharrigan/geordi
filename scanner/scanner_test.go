package scanner_test

import (
	"bytes"
	"github.com/daneharrigan/geordi/scanner"
	"testing"
)

func TestScan(t *testing.T) {
	values := []struct {
		b []byte
		v [][]byte
	}{
		{[]byte("CMD"),           [][]byte{[]byte("CMD")}},
		{[]byte(`CMD "key"`),     [][]byte{[]byte("CMD"), []byte("key")}},
		{[]byte(`CMD "key" 1`),   [][]byte{[]byte("CMD"), []byte("key"), []byte("1")}},
		{[]byte(`CMD "key" 1.0`), [][]byte{[]byte("CMD"), []byte("key"), []byte("1.0")}},
		{[]byte(`CMD "key" -1`),  [][]byte{[]byte("CMD"), []byte("key"), []byte("-1")}},
		{[]byte(`CMD "key" "value"`), [][]byte{[]byte("CMD"), []byte("key"), []byte("value")}},
		{[]byte(`CMD "key" "value" 1 2.0 -3`), [][]byte{
			[]byte("key"), []byte("value"), []byte("1"), []byte("2.0"), []byte("-3")},
		},
	}

	for i, tt := range values {
		s := scanner.New(tt.b)
		for n := 0; n < len(tt.v); n++ {
			incomplete := s.Scan()
			if n+1 == len(tt.v) && incomplete {
				t.Fatalf("at %d:%d expected %v; was %v", i, n, false, incomplete)
			}

			if !bytes.Equal(s.Bytes(), tt.v[n]) {
				t.Fatalf("at %d:%d expected %q; was %q", i, n, tt.v[n], s.Bytes())
			}

			if s.Err() != nil {
				t.Fatalf("at %d expected %v; was %q", i, nil, s.Err())
			}
		}
	}
}

/*
func TestScanErrCmd(t *testing.T) {
	values := []struct {
		b []byte
		e error
	}{
		{[]byte("CMD1"), scanner.ErrUnexpectedCmdChr},
		{[]byte("cmd"), scanner.ErrUnexpectedCmdChr},
		{[]byte("!CMD"), scanner.ErrUnexpectedCmdChr},
		{[]byte(`"CMD"`), scanner.ErrUnexpectedCmdChr},

		{[]byte("CMD foo\r"), scanner.ErrUnexpectedCR},
		{[]byte("\rCMD foo"), scanner.ErrUnexpectedCR},
		{[]byte("CMD\rfoo"), scanner.ErrUnexpectedCR},

		{[]byte(`CMD "`), scanner.ErrArgumentQuoting},
		{[]byte(`CMD "key\"`), scanner.ErrArgumentQuoting},
		{[]byte("CMD key"), scanner.ErrArgumentQuoting},

		{[]byte("\\CMD"), scanner.ErrUnexpectedEscaping},
		{[]byte(`CMD \"key"`), scanner.ErrUnexpectedEscaping},
	}

	for i, tt := range values {
		s := scanner.New(tt.b)
		if s.Scan() {
			t.Fatalf("at %d expected %v; was %v", i, false, true)
		}

		if s.Err() != tt.e {
			t.Fatalf("at %d expected %q; was %q", i, tt.e, s.Err())
		}
	}
}
*/
