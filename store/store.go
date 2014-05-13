package store

import (
	"github.com/daneharrigan/geordi/types"
	"errors"
)

var (
	store = make(map[string]*Record)
	ErrNotFound = errors.New("record not found")
)

type Record struct {
	Value interface{}
	Type types.Type
}

func NewRecord(v interface{}, t types.Type) *Record {
	return &Record{v, t}
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
