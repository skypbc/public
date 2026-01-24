package plugins

import (
	"github.com/skypbc/public/include"
)

type IEngine = include.IEngine
type Settings = include.Settings

type IPlugin = include.IPlugin
type IPluginVars = include.IPluginVars
type IPluginCommands = include.IPluginCommands
type IPluginHooks = include.IPluginHooks
type PluginHookPositionProto = include.PluginHookPositionProto
type IPluginCommandExecutor = include.IPluginCommandExecutor

type IPluginConfig = include.IPluginConfig
type IPluginManager = include.IPluginManager
type IPluginConnect = include.IPluginConnect
type PluginHook = include.PluginHook
type PluginVarItem = include.PluginVarItem

type Kind = include.EventKind
type Category = include.EventCategory
type Action = include.EventAction
type Status = include.EventStatus
