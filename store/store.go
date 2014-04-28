package store

type Type int

type Value interface {
	Type() Type
	Bytes() []byte
}

type value struct {
	dataType Type
	data []byte
}

const (
	String Type = iota
	Int
)

var (
	store = map[string]Value
	InvalidType = errors.New("invalid type")
	NotFound = errors.New("key not found")
)

func (v value) Type() Type {
	return v.dataType
}

func (v value) Bytes() []byte {
	return v.data
}

func Set(key string, val Value) {
	store[key] = val
}

func Get(key string) (Value, error) {
	v, ok := store[key]
	if !ok {
		return nil, NotFound
	}

	return v, nil
}
