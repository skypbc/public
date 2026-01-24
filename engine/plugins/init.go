package plugins

import (
	"context"
)

var NewPluginManager func(configFile string) IPluginManager
var NewPluginVars func(vars ...[]PluginVarItem) []PluginVarItem
var NewConfig func(name string, version string, tag string, active bool, settings Settings) IPluginConfig
var LoadConfig func(ctx context.Context, tag string) (IPluginConfig, error)
var SaveConfig func(ctx context.Context, config IPluginConfig) error
var DeleteConfig func(ctx context.Context, plugin IPlugin) error
var ResetConfigs func(ctx context.Context) (err error)
var SetPluginManager func(ctx context.Context, obj IPluginManager) error
var ExtractPluginManager func(ctx context.Context) IPluginManager
