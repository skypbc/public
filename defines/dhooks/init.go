package dhooks

import (
	"github.com/skypbc/public/engine/hooks"
	"github.com/skypbc/public/include"
)

var HookList = make([]*include.IHook, 0)

func InitHooks() {
	CommandNumber = hooks.NewRawHook(CommandNumberName, CommandNumberHookProto(nil))

	HookList = append(HookList, &CommandNumber)
}
