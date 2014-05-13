package command

import (
	"github.com/daneharrigan/geordi/responder"
	"github.com/daneharrigan/geordi/store"
	"github.com/daneharrigan/geordi/types"
)

func init() {
	commands["GET"] = get
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
