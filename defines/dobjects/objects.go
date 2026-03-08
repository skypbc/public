package dobjects

import (
	"context"

	"github.com/skypbc/goutils/gerrors"
	"github.com/skypbc/public/engine/objects"
)

var (
	None,
	Config,
	PluginManager,
	CommandManager,
	HookManager int
)

var initItems = map[string]*int{
	"none":            &None,
	"config":          &Config,
	"plugin_manager":  &PluginManager,
	"command_manager": &CommandManager,
	"hook_manager":    &HookManager,
}

func Init(ctx context.Context) error {
	objManager := objects.ExtractObjectManager(ctx)
	if objManager == nil {
		return gerrors.NewNotFoundError().
			SetTemplate("ObjectManager not found!")
	}
	for k, v := range initItems {
		index := objManager.Index(k, true)
		if index == -1 {
			return gerrors.NewUnknownError().
				SetTemplate("Failed to create object for key: {key}").
				AddStr("key", k)
		}
		*v = index
	}
	return nil
}
