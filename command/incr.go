package command

import (
	"github.com/daneharrigan/geordi/responder"
	"github.com/daneharrigan/geordi/store"
	"github.com/daneharrigan/geordi/types"
)

func init() {
	commands["INCR"] = incr
}

func incr(respond *responder.Responder, args []Argument) {
	if len(args) < 1 || len(args) > 2 {
		respond.WriteError(ErrArgumentCount)
		return
	}

	if args[0].Type != types.String {
		respond.WriteError(ErrArgumentType)
		return
	}

	if len(args) == 2 && args[1].Type != types.Int &&
		args[1].Type != types.Float {
		respond.WriteError(ErrArgumentType)
		return
	}

	key := string(args[0].Value)
	record, err := store.Get(key)

	if err != nil {
		respond.WriteError(err)
		return
	}

	switch record.Type {
	case types.Int:
		if len(args) < 2 {
			record.Update(record.Int()+1, types.Int)
		} else {
			arg := args[1]
			switch arg.Type {
			case types.Int:
				n := types.ToInt(arg.Value)
				record.Update(record.Int()+n, types.Int)

				respond.SetSuccess()
				respond.WriteInt(record.Int())
			case types.Float:
				n := types.ToFloat(arg.Value)
				record.Update(float64(record.Int())+n, types.Float)

				respond.SetSuccess()
				respond.WriteFloat(record.Float())
			}
		}
	case types.Float:
		if len(args) < 2 {
			record.Update(record.Float()+1, types.Float)
		} else {
			arg := args[1]
			n := types.ToFloat(arg.Value)
			record.Update(record.Float()+n, types.Float)
		}

		respond.SetSuccess()
		respond.WriteFloat(record.Float())
	default:
		respond.WriteError(ErrCmdRecordType)
	}
}
