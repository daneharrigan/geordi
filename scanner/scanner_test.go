package scanner

import (
	"bytes"
	"github.com/daneharrigan/geordi/types"
	"testing"
)

func TestScan(t *testing.T) {
	type value struct {
		b []byte
		t types.Type
	}

	table := []struct {
		operation []byte
		values    []value
		size      int
	}{
		{
			[]byte("CMD"),
			[]value{{[]byte("CMD"), types.String}},
			1,
		},
		{
			[]byte(`CMD "key"`),
			[]value{
				{[]byte("CMD"), types.String},
				{[]byte("key"), types.String},
			},
			2,
		},
		{
			[]byte("CMD 1"),
			[]value{
				{[]byte("CMD"), types.String},
				{[]byte("1"), types.Int},
			},
			2,
		},
		{
			[]byte("CMD .5"),
			[]value{
				{[]byte("CMD"), types.String},
				{[]byte(".5"), types.Float},
			},
			2,
		},
		{
			[]byte("CMD 0.5"),
			[]value{
				{[]byte("CMD"), types.String},
				{[]byte("0.5"), types.Float},
			},
			2,
		},
		{
			[]byte("CMD -1"),
			[]value{
				{[]byte("CMD"), types.String},
				{[]byte("-1"), types.Int},
			},
			2,
		},
		{
			[]byte(`CMD "key" 12`),
			[]value{
				{[]byte("CMD"), types.String},
				{[]byte("key"), types.String},
				{[]byte("12"), types.Int},
			},
			3,
		},
		{
			[]byte(`CMD "key" "value" 1 -1 .5 0.5`),
			[]value{
				{[]byte("CMD"), types.String},
				{[]byte("key"), types.String},
				{[]byte("value"), types.String},
				{[]byte("1"), types.Int},
				{[]byte("-1"), types.Int},
				{[]byte(".5"), types.Float},
				{[]byte("0.5"), types.Float},
			},
			7,
		},

		{
			[]byte(" CMD 1 "),
			[]value{
				{[]byte("CMD"), types.String},
				{[]byte("1"), types.Int},
			},
			2,
		},
		{
			[]byte("\tCMD\t1\t"),
			[]value{
				{[]byte("CMD"), types.String},
				{[]byte("1"), types.Int},
			},
			2,
		},
		{
			[]byte("\nCMD\n1\n"),
			[]value{
				{[]byte("CMD"), types.String},
				{[]byte("1"), types.Int},
			},
			2,
		},
		{
			[]byte("\n\t CMD\n\t 1\n\t "),
			[]value{
				{[]byte("CMD"), types.String},
				{[]byte("1"), types.Int},
			},
			2,
		},
	}

	for _, tt := range table {
		o := tt.operation
		s := New(o)
		var i int

		for n := 0; n < len(tt.values); n++ {
			s.Scan()
			v := tt.values[n]
			a := v.b
			b := s.Bytes()
			i++

			if !bytes.Equal(a, b) {
				t.Fatalf("from %q, expected %q; was %q", o, a, b)
			}

			if s.Type() != v.t {
				t.Fatalf("from %q, expected %q; was %q", o,
					s.Type(), v.t)
			}
		}

		if s.Err() != nil {
			t.Fatalf("expected %q; was nil", s.Err())
		}

		if i < tt.size {
			t.Fatalf("expected %d; was %d", tt.size, i)
		}

	}
}

func TestError(t *testing.T) {
	values := []struct {
		operation []byte
		err       error
	}{
		{[]byte(`"CMD"`), ErrUnexpectedCmdChr},
		{[]byte(`"CMD1`), ErrUnexpectedCmdChr},
		{[]byte(`"CMD!`), ErrUnexpectedCmdChr},
		{[]byte("CMD \r"), ErrUnexpectedCR},
		{[]byte(`CMD key`), ErrArgumentQuoting},
		{[]byte(`CMD "`), ErrArgumentQuoting},
		{[]byte(`CMD \"`), ErrUnexpectedEscaping},
		{[]byte("CMD 1.1.1"), ErrInvalidDecimal},
	}

	for _, tt := range values {
		s := New(tt.operation)
		for s.Scan() {
		}
		if s.Err() != tt.err {
			t.Fatalf("expected %q; was %e", tt.err, s.Err())
		}
	}
}
