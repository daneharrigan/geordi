package command

import (
	"errors"
	"github.com/daneharrigan/geordi/responder"
	"github.com/daneharrigan/geordi/scanner"
	"github.com/daneharrigan/geordi/types"
)

type Argument struct {
	Value []byte
	Type  types.Type
}

type Command func(*responder.Responder, []Argument)

var (
	commands         = make(map[string]Command)
	ErrNotFound      = errors.New("command not found")
	ErrArgumentCount = errors.New("invalid argument count")
	ErrArgumentType  = errors.New("invalid argument type")
	OK               = "OK"
)

func Execute(operation []byte, respond *responder.Responder) error {
	s := scanner.New(operation)
	s.Scan()

	if s.Err() != nil {
		respond.WriteError(s.Err())
		return nil
	}

	idx := string(s.Bytes())
	fn, ok := commands[idx]

	if !ok {
		respond.WriteError(ErrNotFound)
		return nil
	}

	fn(respond, arguments(s))
	return nil
}

func arguments(s *scanner.Scanner) (args []Argument) {
	for s.Scan() {
		args = append(args, Argument{s.Bytes(), s.Type()})
	}

	return
}
