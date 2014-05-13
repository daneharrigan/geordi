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

	b := record.Value.([]byte)
	respond.SetSuccess()

	if record.Type == types.String {
		respond.WriteString(string(b))
	} else {
		respond.Write(b)
	}
}
