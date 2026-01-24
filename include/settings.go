package include

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/skypbc/goutils/gmap"
	"github.com/skypbc/goutils/gnum"
	"github.com/skypbc/public/defines/derrs"
	"strings"
)

type Settings map[string]any

func (s *Settings) Scan(value any) error {
	if value == nil {
		*s = map[string]any{}
		return nil
	}

	var data []byte

	switch x := value.(type) {
	case string:
		// Не знаю почему, но строка приходит с экранированными бэкслешами.
		if strings.Contains(x, `\\`) {
			// Убираем лишние бэкслеши
			x = strings.ReplaceAll(x, `\\`, `\`)
		}
		data = []byte(x)
	case []byte:
		data = x
	default:
		return derrs.NewIncorrectParamsError()
	}

	if err := json.Unmarshal(data, s); err != nil {
		return err
	}

	return nil
}

func (s Settings) String() string {
	if s == nil {
		return "{}"
	}
	data, _ := json.Marshal(s)
	return string(data)
}

func (s Settings) DbIsJsonData() bool {
	return true
}

func (s Settings) Value() (driver.Value, error) {
	if s == nil {
		return []byte("{}"), nil
	}
	return json.Marshal(s)
}

func (s Settings) Get(key string) (val any, ok bool) {
	return gmap.GetAny(s, key)
}

func (s Settings) GetBool(key string) (val bool, ok bool) {
	v, ok := gmap.GetAny(s, key)
	if !ok {
		return false, false
	}
	if val, ok = v.(bool); ok {
		return val, ok
	}
	if num, ok := gnum.TryInt(v); ok {
		return num != 0, ok
	}
	return false, false
}

func (s Settings) GetBoolOrDefault(key string, def bool) bool {
	val, ok := s.GetBool(key)
	if !ok {
		val = def
	}
	return val
}

func (s Settings) GetInt(key string) (val int64, ok bool) {
	v, ok := gmap.GetAny(s, key)
	if !ok {
		return 0, false
	}
	return gnum.TryInt(v)
}

func (s Settings) GetIntOrDefault(key string, def int64) int64 {
	val, ok := s.GetInt(key)
	if !ok {
		val = def
	}
	return val
}

func (s Settings) GetUint(key string) (val uint64, ok bool) {
	v, ok := gmap.GetAny(s, key)
	if !ok {
		return 0, false
	}
	return gnum.TryUint(v)
}

func (s Settings) GetUintOrDefault(key string, def uint64) uint64 {
	val, ok := s.GetUint(key)
	if !ok {
		val = def
	}
	return val
}

func (s Settings) GetFloat(key string) (val float64, ok bool) {
	v, ok := gmap.GetAny(s, key)
	if !ok {
		return 0, false
	}
	return gnum.TryFloat(v)
}

func (s Settings) GetFloatOrDefault(key string, def float64) float64 {
	val, ok := s.GetFloat(key)
	if !ok {
		val = def
	}
	return val
}

func (s Settings) GetString(key string) (val string, ok bool) {
	v, ok := gmap.GetAny(s, key)
	if !ok {
		return "", false
	}
	val, ok = v.(string)
	return val, ok
}

func (s Settings) GetStringOrDefault(key string, def string) string {
	val, ok := s.GetString(key)
	if !ok {
		val = def
	}
	return val
}

func (s Settings) Set(key string, val any) {
	s[key] = val
}
