package command

import (
	"errors"
	"github.com/daneharrigan/geordi/responder"
	"github.com/daneharrigan/geordi/scanner"
	"github.com/daneharrigan/geordi/store"
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

func init() {
	commands["GET"] = get
	commands["SET"] = set
}

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

func get(respond *responder.Responder, args []Argument) {
	if len(args) != 1 {
		respond.WriteError(ErrArgumentCount)
		return
	}

	if args[0].Type != types.String {
		respond.WriteError(ErrArgumentType)
		return
	}

	key := string(args[0].Value)
	record, err := store.Get(key)

	if err != nil {
		respond.WriteError(err)
		return
	}

	b := record.Value.([]byte)
	respond.SetSuccess()

	if record.Type == types.String {
		respond.WriteString(string(b))
	} else {
		respond.Write(b)
	}
}

func set(respond *responder.Responder, args []Argument) {
	if len(args) != 2 {
		respond.WriteError(ErrArgumentCount)
		return
	}

	key := string(args[0].Value)
	value := args[1]
	record := store.NewRecord(value.Value, value.Type)
	store.Set(key, record)

	respond.SetSuccess()
	respond.WriteString(OK)
}
