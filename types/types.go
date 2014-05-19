package types

import "strconv"

type Type int

const (
	String Type = iota
	Int
	Float
	Hash
	List
)

func ToInt(b []byte) int64 {
	i, err := strconv.ParseInt(string(b), 0, 64)
	if err != nil {
		panic(err)
	}

	return i
}

func ToFloat(b []byte) float64 {
	f, err := strconv.ParseFloat(string(b), 64)
	if err != nil {
		panic(err)
	}

	return f
}
