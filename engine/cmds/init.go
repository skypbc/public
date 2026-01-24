package cmds

import (
	"context"
)

var NewCommandManager func() (res ICommandManager)
var SetCommandManager func(ctx context.Context, manager ICommandManager) error
var ExtractCommandManager func(ctx context.Context) ICommandManager

var NewCommandExecutor func(plugin IPlugin) (ICommandExecutor, error)
var NewCommandInfo func(plugin IPlugin, cmd ICommand, hooks ...ICommandHooks) ICommandInfo

var NewCommandSignature func(kind Kind, category []Category, action Action) ICommandSignature
var NewCommandHooks func() ICommandHooks
