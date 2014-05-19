package store

import (
	"errors"
	"github.com/daneharrigan/geordi/types"
)

var (
	store       = make(map[string]*Record)
	ErrNotFound = errors.New("record not found")
)

type Record struct {
	Value interface{}
	Type  types.Type
}

func NewRecord(b []byte, t types.Type) *Record {
	r := &Record{Type: t}
	switch t {
	case types.Int:
		r.Value = types.ToInt(b)
	case types.Float:
		r.Value = types.ToFloat(b)
	default:
		r.Value = b
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

func (r *Record) Update(v interface{}, t types.Type) {
	r.Type = t
	r.Value = v
}
