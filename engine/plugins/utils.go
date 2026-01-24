package plugins

import (
	"fmt"
	"github.com/skypbc/goutils/greflect"
	"github.com/skypbc/public/defines/derrs"
	"plugin"
	"reflect"
)

// FIX_ME: Когда появятся алиасы, переместить реализацию в "private/engine/plugins"
func Lookup[T any](plugin *plugin.Plugin, name string) (val *T, err error) {
	symbol, err := plugin.Lookup(name)
	if err != nil {
		return nil, derrs.NewPluginError(err).
			SetTemplate(`failed to find symbol "%s"`).
			AddStr("name", name)
	}
	var rv reflect.Value
	for ok := true; ok; {
		err = nil
		switch v := symbol.(type) {
		case T:
			val = &v
		case *T:
			val = v
		default:
			err = derrs.NewPluginError().
				SetTemplate(`invalid symbol type "%T"`).
				AddStr("symbol", fmt.Sprintf("%T", symbol))
		}
		if val != nil {
			break
		}
		if rv == (reflect.Value{}) {
			rv = reflect.ValueOf(symbol)
		}
		rv, ok = greflect.TryElem(rv)
		symbol = rv.Interface()
	}
	return val, err
}
