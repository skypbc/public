package dconfig

import (
	"context"
	"github.com/skypbc/public/engine/config"
	"github.com/skypbc/public/include"
)

var initItems = []include.ConfigInitItem{
	{
		Key: SKY_EVENT_MAX_META_SIZE_S, Index: &SKY_EVENT_MAX_META_SIZE,
		DefaultValue: 16 * 1024, CacheItemRef: &cacheEventMaxMetaSize,
	},

	{
		Key:          SKY_EVENT_DATA_FUNC_S,
		Index:        &SKY_EVENT_DATA_FUNC,
		CacheItemRef: &cacheEventDataFunc,
	}, {
		Key:          SKY_EVENT_SET_DATA_FUNC_S,
		Index:        &SKY_EVENT_SET_DATA_FUNC,
		CacheItemRef: &cacheEventSetDataFunc,
	},

	{
		Key:          SKY_EVENT_SET_DATA_RAW_FUNC_S,
		Index:        &SKY_EVENT_SET_DATA_RAW_FUNC,
		CacheItemRef: &cacheEventSetDataRawFunc,
	},
}

func Init(ctx context.Context) error {
	return config.InitCacheItems(ctx, initItems)
}
