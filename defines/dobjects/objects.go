package dobjects

import (
	"context"
	"github.com/skypbc/public/engine/config"
	"github.com/skypbc/public/include"
)

var (
	None,
	Config,
	PluginManager,
	CommandManager,
	HookManager int
)

var initItems = []include.ConfigInitItem{
	{Key: "none", ItemRef: &None},

	{Key: "config", ItemRef: &Config},
	{Key: "plugin_manager", ItemRef: &PluginManager},
	{Key: "command_manager", ItemRef: &CommandManager},
	{Key: "hook_manager", ItemRef: &HookManager},
}

func Init(ctx context.Context) error {
	return config.InitItems(ctx, initItems, "defines.objects")
}
