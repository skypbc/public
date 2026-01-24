package include

import "context"

type HookExecuteProto func(ctx context.Context, hookImpl any) (next bool, err error)

type IHook interface {
	String() string
	Index() int

	Name() string
	Prototype() any

	Add(plugin IPlugin, hookImpl any, posX int, posY int) error

	PreImplList() []IHookImpl
	ImplList() []IHookImpl
	PostImplList() []IHookImpl

	DisconnectPlugin(plugin IPlugin)
}

type IHookImpl interface {
	Get() any
	Plugin() IPlugin
}

type IHookManager interface {
	IPluginConnect

	Register(hookName string, prototype any) (IHook, error)

	Get(name string) IHook

	Execute(ctx context.Context, hookName string, f HookExecuteProto) (err error)
	Execute2(ctx context.Context, hook IHook, f HookExecuteProto) (err error)
	Execute3(ctx context.Context, index int, f HookExecuteProto) (err error)

	MustExecute(ctx context.Context, hookName string, f HookExecuteProto) (err error)
	MustExecute2(ctx context.Context, hook IHook, f HookExecuteProto) (err error)
	MustExecute3(ctx context.Context, index int, f HookExecuteProto) (err error)
}
