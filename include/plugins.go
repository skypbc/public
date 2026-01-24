package include

import (
	"context"
)

type PluginHookPositionProto func(ctx context.Context, pluginHook PluginHook, hook IHook) (int, error)

// FIX_ME: Сделать:
// PluginHookOpts из PluginHook
// Добавить метод NewPluginHook(opts PluginHookOpts) IPluginHook
// Таким образом хук можно будет определить через вызов метода NewPluginHook или при помощи определения
// полноценной структуры поддерживающей интерфейс IPluginHook.

// При создании PluginHook, необходимо задать либо RawHook, либо Name + Prototype.
type PluginHook struct {
	RawHook *IHook
	// Name указывается если не задан RawHook
	Name string
	// Prototype указывается если не задан RawHook. При этом Prototype может отсуствовать, если известно,
	// что хук уже зарегистрирован.
	Prototype any
	// Optional. Реализация может отсуствовать, если вы хотите только получить ссылку на хук.
	Impl any
	// Optional. Определяет позицию, в которую нужно поместить реализацию.
	Position int
	// Optional. Выполняется перед хуком
	PreHook bool
	// Optional. Выполняется после хука
	PostHook bool
	// Optional. Колбек используется для определения позиции, в которую нужно поместить реализацию.
	// Если он не задн, используется pluginHook.Position
	PositionCallback PluginHookPositionProto
}

type IPluginManager interface {
	Get(tag string) IPlugin
	Config(ctx context.Context, tag string) IPluginConfig
	Register(ctx context.Context, plugin IPlugin, external bool) (new bool, err error)
	Load(ctx context.Context, plugin IPlugin) error

	MainConfig() IPluginMainConfig
	ReloadMainConfig() error

	DeleteConfig(ctx context.Context, plugin IPlugin) error
	ResetConfigs(ctx context.Context) error
}

type IPluginMainConfig interface {
	Plugins() []PluginMainConfigPluginItem
}

type PluginMainConfigPluginItem struct {
	Tag    string `json:"tag"`
	Active bool   `json:"active"`
	Source string `json:"source"`
	Path   string `json:"path"`
}

type IPlugin interface {
	Name() string
	Version() string
	Tag() string

	Active() bool
	Settings() Settings

	PreInit(ctx context.Context, install bool) error
	Init(ctx context.Context, install bool) error
	Update(ctx context.Context, pluginConfig IPluginConfig) error
	Load(ctx context.Context, pluginConfig IPluginConfig) error
}

type IPluginVars interface {
	Vars() []PluginVarItem
}

type IPluginCommands interface {
	Commands() []ICommand
}

type IPluginHooks interface {
	Hooks() []PluginHook
}

type IPluginConfigInitItems interface {
	ConfigInitItems() []ConfigInitItem
}

type IPluginCommandExecutor interface {
	CommandExecutor() CommandExecutorProto
}

type IPluginConfig interface {
	Name() string
	Version() string
	Tag() string
	Active() bool
	Settings() Settings

	Update(plugin IPlugin)
}

type IPluginConnect interface {
	ConnectPlugin(ctx context.Context, plugin IPlugin) error
	DisconnectPlugin(ctx context.Context, plugin IPlugin) error
}

type PluginVarItem struct {
	Group string
	Name  string
	// Значения зависят от модуля использующего VarItem.
	// В Value обычно указывается указатель на переменную, которая будет хранить Value/Number/etc.
	Value        any
	DefaultValue any
}
