package command

import (
	"bufio"
	"errors"
	"github.com/daneharrigan/geordi/store"
	"io"
)

type Handler func([][]byte) ([]byte, error)

type Command interface {
	Run() ([]byte, error)
}

type command struct {
	handler Handler
	args    [][]byte
}

var (
	NotFound    = errors.New("command not found")
	InvalidArgs = errors.New("invalid arguments")
	commands    = make(map[string]Handler)
)

func init() {
	commands["SET"] = set
	commands["GET"] = get
	commands["KEYS"] = keys
}

func Parse(r io.Reader) (Command, error) {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)

	name := s.Text()
	fn, ok := commands[name]
	if !ok {
		return nil, NotFound
	}

	var args [][]byte
	for s.Scan() {
		args = append(args, s.Bytes())
	}

	return command{fn, args}, nil
}

func (c command) Run() ([]byte, error) {
	return c.handler(c.args)
}

func set(args [][]byte) ([]byte, error) {
	if len(args) != 2 {
		return nil, InvalidArgs
	}

	k := string(args[0])
	v := store.NewValue(args[1])
	return store.Set(k, v)
}

func get(args [][]byte) ([]byte, error) {
	if len(args) != 1 {
		return nil, InvalidArgs
	}

	return store.Get(string(args[0]))
}

func keys(_ [][]byte) ([]byte, error) {
	return store.Keys()
}
