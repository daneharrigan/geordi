package command

import (
	"github.com/daneharrigan/geordi/responder"
	"github.com/daneharrigan/geordi/store"
	"github.com/daneharrigan/geordi/types"
)

func init() {
	commands["HSET"] = hset
}

func hset(respond *responder.Responder, args []Argument) {
	if len(args) != 3 {
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

	if err == store.ErrNotFound {
		record = store.NewHash()
		store.Set(key, record)
	} else if err != nil {
		respond.WriteError(err)
		return
	}

	hash := record.Hash()
	key = string(args[1].Value)
	value := args[2]
	record = store.NewRecord(value.Value, value.Type)

	hash[key] = record
	respond.SetSuccess()
	respond.WriteString(OK)
}
