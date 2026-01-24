package dhooks

import (
	"context"
)

var CommandNumber IHook
var CommandNumberName = "ENGINE_COMMAND_NUMBER"

type CommandNumberHookProto func(ctx context.Context, sig ICommandSignature) (number int64, next bool, err error)
