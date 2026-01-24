package include

import (
	"github.com/skypbc/goutils/gfmt"
	"github.com/skypbc/public/defines/derrs"
	"reflect"
	"sync"
)

type ConfigCacheItem[T any] struct {
	item    *ConfigItem
	value   T
	updated int64
	setLock *sync.Mutex
}

func NewConfigCacheItem[T any]() ConfigCacheItem[T] {
	return ConfigCacheItem[T]{
		item:    nil,
		updated: 0,
		setLock: &sync.Mutex{},
	}
}

func (c *ConfigCacheItem[T]) Set(item *ConfigItem) bool {
	c.setLock.Lock()
	defer c.setLock.Unlock()

	if item.value == nil {
		c.item = item
		return true
	}

	val, updated := item.Value2()
	rSrc := reflect.ValueOf(val)
	rDst := reflect.ValueOf(&c.value).Elem()

	if rSrc.CanConvert(rDst.Type()) {
		c.updated = updated
		c.item = item
		rDst.Set(rSrc.Convert(rDst.Type()))
		return true
	}
	return false
}

func (c *ConfigCacheItem[T]) Key() string {
	return c.item.key
}

func (c *ConfigCacheItem[T]) Value() T {
	if c.updated == c.item.updated {
		return c.value
	}
	if !c.Set(c.item) {
		err := derrs.NewConfigError().
			SetTemplate("ConfigCacheItem: Key \"{key}\" has incompatible value \"{value}\" with type {type}.").
			AddStr("key", c.item.Key()).
			AddAny("value", gfmt.Sprintf("%v", c.item.Value())).
			AddAny("type", gfmt.Sprintf("%T", c.value))
		panic(err)
	}
	return c.value
}

func (c *ConfigCacheItem[T]) Raw() *ConfigItem {
	return c.item
}
