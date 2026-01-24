package include

import (
	"github.com/skypbc/goutils/gfmt"
	"github.com/skypbc/goutils/gnum"
	"github.com/skypbc/goutils/greflect"
	"github.com/skypbc/goutils/gstring"
	"github.com/skypbc/public/defines/derrs"
	"sync"
)

func NewConfigItem(index int, key string, val any) *ConfigItem {
	return &ConfigItem{
		key:     key,
		index:   index,
		value:   val,
		updated: 1,
		setLock: &sync.RWMutex{},
	}
}

type ConfigItem struct {
	key     string
	index   int
	value   any
	updated int64
	setLock *sync.RWMutex
}

func (c *ConfigItem) Key() string {
	return c.key
}

func (c *ConfigItem) Index() int {
	return c.index
}

func (c *ConfigItem) SetValue(v any) {
	c.setLock.Lock()
	defer c.setLock.Unlock()
	c.value = v
	c.updated++
}

func (c *ConfigItem) Value() any {
	return c.value
}

func (c *ConfigItem) Value2() (val any, updated int64) {
	c.setLock.RLock()
	defer c.setLock.RUnlock()
	return c.value, c.updated
}

func (c *ConfigItem) Updated() int64 {
	return c.updated
}

func (c *ConfigItem) Fill(dest any) {
	if ok := c.TryFill(dest); !ok {
		err := derrs.NewConfigError().
			SetTemplate("Key: \"{key}\", Val: \"{value}\", Dest: \"{dest}\" can't fill.").
			AddStr("key", c.key).
			AddAny("value", gfmt.Sprintf("%#v", c.value)).
			AddAny("dest", gfmt.Sprintf("%#v", dest))

		panic(err)
	}
}

func (c *ConfigItem) String() string {
	if res, ok := c.TryString(); ok {
		return res
	}
	err := derrs.NewConfigError().
		SetTemplate("Key: \"{key}\", Val: \"{value}\" isn't string.").
		AddStr("key", c.key).
		AddAny("value", gfmt.Sprintf("%#v", c.value))
	panic(err)
}

func (c *ConfigItem) Int64() int64 {
	if v, ok := c.TryInt64(); ok {
		return v
	}
	err := derrs.NewConfigError().
		SetTemplate("Key: \"{key}\", Val: \"{value}\" isn't int64.").
		AddStr("key", c.key).
		AddAny("value", gfmt.Sprintf("%#v", c.value))
	panic(err)
}

func (c *ConfigItem) Int() int {
	if val, ok := c.TryInt64(); ok {
		return int(val)
	}
	err := derrs.NewConfigError().
		SetTemplate("Key: \"{key}\", Val: \"{value}\" isn't int.").
		AddStr("key", c.key).
		AddAny("value", gfmt.Sprintf("%#v", c.value))
	panic(err)
}

func (c *ConfigItem) Int8() int8 {
	if val, ok := c.TryInt64(); ok {
		return int8(val)
	}
	err := derrs.NewConfigError().
		SetTemplate("Key: \"{key}\", Val: \"{value}\" isn't int8.").
		AddStr("key", c.key).
		AddAny("value", gfmt.Sprintf("%#v", c.value))
	panic(err)
}

func (c *ConfigItem) Int16() int16 {
	if val, ok := c.TryInt64(); ok {
		return int16(val)
	}
	err := derrs.NewConfigError().
		SetTemplate("Key: \"{key}\", Val: \"{value}\" isn't int16.").
		AddStr("key", c.key).
		AddAny("value", gfmt.Sprintf("%#v", c.value))
	panic(err)
}

func (c *ConfigItem) Int32() int32 {
	if val, ok := c.TryInt64(); ok {
		return int32(val)
	}
	err := derrs.NewConfigError().
		SetTemplate("Key: \"{key}\", Val: \"{value}\" isn't int32.").
		AddStr("key", c.key).
		AddAny("value", gfmt.Sprintf("%#v", c.value))
	panic(err)
}

func (c *ConfigItem) Uint64() uint64 {
	if v, ok := gnum.TryUint(c.value); ok {
		return v
	}
	err := derrs.NewConfigError().
		SetTemplate("Key: \"{key}\", Val: \"{value}\" isn't uint64.").
		AddStr("key", c.key).
		AddAny("value", gfmt.Sprintf("%#v", c.value))
	panic(err)
}

func (c *ConfigItem) Uint() uint {
	if val, ok := c.TryUint64(); ok {
		return uint(val)
	}
	err := derrs.NewConfigError().
		SetTemplate("Key: \"{key}\", Val: \"{value}\" isn't uint.").
		AddStr("key", c.key).
		AddAny("value", gfmt.Sprintf("%#v", c.value))
	panic(err)
}

func (c *ConfigItem) Uint8() uint8 {
	if val, ok := c.TryUint64(); ok {
		return uint8(val)
	}
	err := derrs.NewConfigError().
		SetTemplate("Key: \"{key}\", Val: \"{value}\" isn't uint8.").
		AddStr("key", c.key).
		AddAny("value", gfmt.Sprintf("%#v", c.value))
	panic(err)
}

func (c *ConfigItem) Uint16() uint16 {
	if val, ok := c.TryUint64(); ok {
		return uint16(val)
	}
	err := derrs.NewConfigError().
		SetTemplate("Key: \"{key}\", Val: \"{value}\" isn't uint16.").
		AddStr("key", c.key).
		AddAny("value", gfmt.Sprintf("%#v", c.value))
	panic(err)
}

func (c *ConfigItem) Uint32() uint32 {
	if val, ok := c.TryUint64(); ok {
		return uint32(val)
	}
	err := derrs.NewConfigError().
		SetTemplate("Key: \"{key}\", Val: \"{value}\" isn't uint32.").
		AddStr("key", c.key).
		AddAny("value", gfmt.Sprintf("%#v", c.value))
	panic(err)
}

func (c *ConfigItem) Float64() float64 {
	if v, ok := gnum.TryFloat(c.value); ok {
		return v
	}
	err := derrs.NewConfigError().
		SetTemplate("Key: \"{key}\", Val: \"{value}\" isn't float64.").
		AddStr("key", c.key).
		AddAny("value", gfmt.Sprintf("%#v", c.value))
	panic(err)
}

func (c *ConfigItem) Float32() float32 {
	if val, ok := c.TryFloat64(); ok {
		return float32(val)
	}
	err := derrs.NewConfigError().
		SetTemplate("Key: \"{key}\", Val: \"{value}\" isn't float32.").
		AddStr("key", c.key).
		AddAny("value", gfmt.Sprintf("%#v", c.value))
	panic(err)
}

func (c *ConfigItem) Bool() bool {
	if v, ok := gnum.TryBool(c.value); ok {
		return v
	}
	err := derrs.NewConfigError().
		SetTemplate("Key: \"{key}\", Val: \"{value}\" isn't bool.").
		AddStr("key", c.key).
		AddAny("value", gfmt.Sprintf("%#v", c.value))
	panic(err)
}

func (c *ConfigItem) TryFill(dest any) (ok bool) {
	if c.value == nil {
		err := derrs.NewConfigError().
			SetTemplate("Key \"{key}\" not found.").
			AddStr("key", c.key)
		panic(err)
	}
	return greflect.TryFill(c.value, dest)
}

func (c *ConfigItem) TryString() (res string, ok bool) {
	return gstring.TryString(c.value)
}

func (c *ConfigItem) TryInt64() (int64, bool) {
	if v, ok := gnum.TryInt(c.value); ok {
		return v, ok
	}
	return 0, false
}

func (c *ConfigItem) TryInt() (int, bool) {
	if val, ok := c.TryInt64(); ok {
		return int(val), true
	}
	return 0, false
}

func (c *ConfigItem) TryInt8() (int8, bool) {
	if val, ok := c.TryInt64(); ok {
		return int8(val), true
	}
	return 0, false
}

func (c *ConfigItem) TryInt16() (int16, bool) {
	if val, ok := c.TryInt64(); ok {
		return int16(val), true
	}
	return 0, false
}

func (c *ConfigItem) TryInt32() (int32, bool) {
	if val, ok := c.TryInt64(); ok {
		return int32(val), true
	}
	return 0, false
}

func (c *ConfigItem) TryUint64() (uint64, bool) {
	if v, ok := gnum.TryUint(c.value); ok {
		return v, ok
	}
	return 0, false
}

func (c *ConfigItem) TryUint() (uint, bool) {
	if val, ok := c.TryUint64(); ok {
		return uint(val), true
	}
	return 0, false
}

func (c *ConfigItem) TryUint8() (uint8, bool) {
	if val, ok := c.TryUint64(); ok {
		return uint8(val), true
	}
	return 0, false
}

func (c *ConfigItem) TryUint16() (uint16, bool) {
	if val, ok := c.TryUint64(); ok {
		return uint16(val), true
	}
	return 0, false
}

func (c *ConfigItem) TryUint32() (uint32, bool) {
	if val, ok := c.TryUint64(); ok {
		return uint32(val), true
	}
	return 0, false
}

func (c *ConfigItem) TryFloat64() (float64, bool) {
	if v, ok := gnum.TryFloat(c.value); ok {
		return v, true
	}
	return 0, false
}

func (c *ConfigItem) TryFloat32() (float32, bool) {
	if val, ok := c.TryFloat64(); ok {
		return float32(val), true
	}
	return 0, false
}

func (c *ConfigItem) TryBool() (bool, bool) {
	if v, ok := gnum.TryBool(c.value); ok {
		return v, true
	}
	return false, false
}
