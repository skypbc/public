package dconfig

import (
	"context"
	"github.com/skypbc/public/include"
)

var (
	cacheEventMaxMetaSize = include.NewConfigCacheItem[int64]()

	cacheEventDataFunc    = include.NewConfigCacheItem[EventDataProto]()
	cacheEventSetDataFunc = include.NewConfigCacheItem[EventSetDataProto]()

	cacheEventSetDataRawFunc = include.NewConfigCacheItem[EventSetDataProto]()
)

func EventMaxMetaSize() int64 {
	return cacheEventMaxMetaSize.Value()
}

func EventData(ctx context.Context, event IEvent) (data IData, err error) {
	fn := cacheEventDataFunc.Value()
	return fn(ctx, event)
}

func EventSetData(ctx context.Context, event IEvent, data any, size ...int64) (err error) {
	fn := cacheEventSetDataFunc.Value()
	return fn(ctx, event, data, size...)
}

func EventSetDataFunc(cfg IEngineConfig, fn EventDataProto) error {
	return cfg.Set(SKY_EVENT_DATA_FUNC, fn)
}

func EventSetEventSetDataFunc(cfg IEngineConfig, fn EventSetDataProto) error {
	return cfg.Set(SKY_EVENT_SET_DATA_FUNC, fn)
}

func EventSetDataRaw(ctx context.Context, event IEvent, data any, size ...int64) (err error) {
	fn := cacheEventSetDataRawFunc.Value()
	return fn(ctx, event, data, size...)
}

func EventSetEventSetDataRawFunc(cfg IEngineConfig, fn EventSetDataProto) error {
	return cfg.Set(SKY_EVENT_SET_DATA_RAW_FUNC, fn)
}
