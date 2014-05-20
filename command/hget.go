package command

import (
	"github.com/daneharrigan/geordi/responder"
	"github.com/daneharrigan/geordi/store"
	"github.com/daneharrigan/geordi/types"
)

func init() {
	commands["HGET"] = hget
}

func hget(respond *responder.Responder, args []Argument) {
	if len(args) != 2 {
		respond.WriteError(ErrArgumentCount)
		return
	}

	if args[0].Type != types.String {
		respond.WriteError(ErrArgumentType)
		return
	}

	if args[1].Type != types.String {
		respond.WriteError(ErrArgumentType)
		return
	}

	key := string(args[0].Value)
	record, err := store.Get(key)

	if err != nil {
		respond.WriteError(err)
		return
	}

	key = string(args[1].Value)
	hash := record.Hash()
	record, ok := hash[key]

	if !ok {
		respond.WriteError(store.ErrHashValueNotFound)
		return
	}

	respond.SetSuccess()

	switch record.Type {
	case types.String:
		respond.WriteString(record.String())
	case types.Int:
		respond.WriteInt(record.Int())
	case types.Float:
		respond.WriteFloat(record.Float())
	}
}
