package command

/*

import (
	"bufio"
	"io"
	"errors"
)

const (
	String kind = iota
	Int
)

type kind int

type argument struct {
	kind kind
	value interface{}
}

type command struct {
	name string
	arguments []argument
}

type Command interface {
	Run() (string, error)
}

type Handler func([]string) (string, error)

var (
	commands map[string]Handler
	NotFound error
	InvalidArguments error
)

func init() {
	NotFound = errors.New("command not found")
	InvalidArguments = errors.New("invalid arguments")

	commands = make(map[string]Handler)
	commands["SET"] = handlerSet
	commands["GET"] = handlerGet
}

func Parse(r io.Reader) (c Command, err error){
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)
	if s.Scan() {
		c.name = s.Text()
	}

	for s.Scan() {
		str := s.Text()
		var v interface{}
		var k kind

		if str[:1] == `"` {
			v, err = strconv.Unquote(str)
			k = String
		} else {
			var n int64
			v, err = strconv.ParseInt(v, 0, 64)
			k = Int
		}

		if err != nil {
			return
		}

		c.arguments = append(c.arguments, argument{k, v})
	}

	return
}

func (c command) Run() (string, error) {
	fn, ok := commands[c.name]
	if !ok {
		return "", NotFound

	}
	return fn(c.arguments)
}

func handlerSet(args []string) (string, error) {
	if len(args) != 2 {
		return "", InvalidArguments
	}

	return store.SetString(args[0], args[1])
}

func handlerGet(args []string) (string, error) {
	if len(args) != 1 {
		return "", InvalidArguments
	}

	return store.GetString(args[0])
}
*/
