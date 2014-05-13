package command

import (
	"github.com/daneharrigan/geordi/responder"
	"github.com/daneharrigan/geordi/store"
)

func init() {
	commands["SET"] = set
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
