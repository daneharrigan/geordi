package store

import "errors"

type Type int

type Value interface {
	Type() Type
	Bytes() []byte
}

type value struct {
	t Type
	b []byte
}

const (
	String Type = iota
	Int
)

var (
	store       = make(map[string]Value)
	InvalidType = errors.New("invalid type")
	NotFound    = errors.New("key not found")
	OK          = []byte("OK")
)

func (v value) Type() Type {
	return v.t
}

func (v value) Bytes() []byte {
	return v.b
}

func Set(key string, val Value) ([]byte, error) {
	store[key] = val
	return OK, nil
}

func Get(key string) ([]byte, error) {
	v, ok := store[key]
	if !ok {
		return nil, NotFound
	}

	return v.Bytes(), nil
}

func Keys() (b []byte, err error) {
	for k, _ := range store {
		if len(b) > 0 {
			b = append(b, '\n')
		}
		b = append(b, []byte(k)...)
	}

	return
}

func NewValue(b []byte) Value {
	return value{detect(b), b}
}

func detect(b []byte) Type {
	for _, c := range b {
		if c <= 48 || c >= 57 {
			return String
		}
	}

	return Int
}
