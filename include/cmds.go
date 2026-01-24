package include

import (
	"context"
	"fmt"
)

type ICommandReplacer interface {
	// ShouldReplace возвращает true, если команда должна быть установлена, внезависимости от того, существует ли
	// другая команда с таким именем или нет.
	ShouldReplace() bool
}

// ICommand интерфейс описывает базовые методы команд
type ICommand interface {
	Name() string

	Kind() EventKind
	Category() []EventCategory
	Action() EventAction

	CheckAccess(ctx context.Context, eventIn IEvent) (ok bool, err error)
	Execute(ctx context.Context, eventIn IEvent) (eventOut IEvent, err error)

	CheckOnResultAccess(ctx context.Context, event IEvent) (ok bool, err error)
	OnResult(ctx context.Context, event IEvent) (IEvent, error)
	OnResultError(ctx context.Context, event IEvent) (IEvent, error)
}

type ICommandSignature interface {
	Kind() EventKind
	Category() []EventCategory
	Action() EventAction
}

type ICommandInfo interface {
	Plugin() IPlugin
	Command() ICommand
	Hooks() ICommandHooks

	ExecutePreHook(ctx context.Context, eventIn IEvent) error
	AddPreHook(plugin IPlugin, fn CmdExecuteHookProtoPre)
	RemovePreHook(plugin IPlugin, fn ...CmdExecuteHookProtoPre)

	ExecutePostHook(ctx context.Context, eventIn IEvent, eventOut IEvent, err error) error
	AddPostHook(plugin IPlugin, fn CmdExecuteHookProtoPost)
	RemovePostHook(plugin IPlugin, fn ...CmdExecuteHookProtoPost)
}

type CommandExecutorProto func(ctx context.Context, cmdInfo ICommandInfo, eventIn IEvent) (eventOut IEvent, err error)

type ICommandExecutor interface {
	Plugin() IPlugin
	Execute(ctx context.Context, cmdInfo ICommandInfo, eventIn IEvent) (eventOut IEvent, err error)
}

type ICommandManager interface {
	fmt.Stringer
	IPluginConnect

	Executor() ICommandExecutor
	SetExecutor(ctx context.Context, plugin IPlugin) (err error)

	Execute(ctx context.Context, event IEvent) (IEvent, error)

	Id(sig ICommandSignature) (cmdId int64, err error)
	Name(sig ICommandSignature) (name string, err error)
	Path(sig ICommandSignature) string
	Signature(cmdId int64) (sig ICommandSignature, err error)

	AddPreHook(plugin IPlugin, fn CmdExecuteHookProtoPre)
	RemovePreHook(plugin IPlugin, fn ...CmdExecuteHookProtoPre)

	AddPostHook(plugin IPlugin, fn CmdExecuteHookProtoPost)
	RemovePostHook(plugin IPlugin, fn ...CmdExecuteHookProtoPost)
}

type CmdExecuteHookProtoPre func(ctx context.Context, cmdInfo ICommandInfo, eventIn IEvent) error
type CmdExecuteHookProtoPost func(ctx context.Context, cmdInfo ICommandInfo, eventIn IEvent, eventOut IEvent, errOut error) error

type ICommandHooks interface {
	ExecutePre(ctx context.Context, cmdInfo ICommandInfo, eventIn IEvent) error
	ExecutePost(ctx context.Context, cmdInfo ICommandInfo, eventIn IEvent, eventOut IEvent, err error) error

	AddPre(plugin IPlugin, fn CmdExecuteHookProtoPre)
	RemovePre(plugin IPlugin, fn ...CmdExecuteHookProtoPre)

	AddPost(plugin IPlugin, fn CmdExecuteHookProtoPost)
	RemovePost(plugin IPlugin, fn ...CmdExecuteHookProtoPost)
}
