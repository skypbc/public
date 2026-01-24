package config

import (
	"context"
)

var NewConfig func(key2value map[string]any, filepath string) (IEngineConfig, error)
var NewConfigItem func(index int, key string, val any) *ConfigItem
var SetConfig func(ctx context.Context, config IEngineConfig) (context.Context, error)
var ExtractConfig func(ctx context.Context) IEngineConfig

var InitItems func(ctx context.Context, items []ConfigInitItem, parentKey ...string) error
var InitCacheItems func(ctx context.Context, items []ConfigInitItem, parentKey ...string) error
var SetCacheItem func(cfg IEngineConfig, cfgItem *ConfigItem, cacheItemRef any) error
