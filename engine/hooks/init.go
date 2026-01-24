package hooks

import (
	"context"
)

var NewHook func(index int, name string, callbackType any) IHook
var NewHookManager func() IHookManager

var NewRawHook func(name string, callbackType any) IHook
var NewHookImpl func(callback any, plugin IPlugin) IHookImpl

var SetHookManager func(ctx context.Context, store IHookManager) error
var ExtractHookManager func(ctx context.Context) IHookManager

var Register func(ctx context.Context, items []*IHook) (err error)
