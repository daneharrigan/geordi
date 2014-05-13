package store

import (
	"errors"
	"github.com/daneharrigan/geordi/types"
	"strconv"
)

var (
	store       = make(map[string]*Record)
	ErrNotFound = errors.New("record not found")
)

type Record struct {
	Value interface{}
	Type  types.Type
}

func NewRecord(v interface{}, t types.Type) *Record {
	r := &Record{Type: t}
	switch t {
	case types.Int:
		s := string(v.([]byte))
		i, err := strconv.ParseInt(s, 0, 64)
		if err != nil {
			panic(err)
		}

		r.Value = i
	case types.Float:
		s := string(v.([]byte))
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			panic(err)
		}

		r.Value = f
	default:
		r.Value = v
	}

	return r
}

func Set(k string, r *Record) {
	store[k] = r
}

func Get(k string) (r *Record, err error) {
	r, ok := store[k]
	if ok {
		return r, nil
	}

	return nil, ErrNotFound
}

func (r *Record) String() string {
	return string(r.Value.([]byte))
}

func (r *Record) Int() int64 {
	return r.Value.(int64)
}

func (r *Record) Float() float64 {
	return r.Value.(float64)
}
