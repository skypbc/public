package include

// Type in [Int,Float,Bool,String,Uuid,Path,time.Second,Any]
type ConfigInitItem struct {
	Key          string
	Index        *int
	ItemRef      any
	DefaultValue any
	Optional     bool
	Type         string
	CacheItemRef any
	// Указывает значение устанавливается программно и может остуствовать в конфигурационном файле
	Internal bool
}

type IEngineConfig interface {
	IPluginConnect

	Filepath() string
	Raw() (raw map[string]any, err error)

	PrepareKeys(key ...string) (key2index map[string]int, err error)

	Index(key string) (index int)
	Key(index int) (key string)
	KeyList() []string

	Item(index int) *ConfigItem
	Item2(key string) *ConfigItem
	Count() int

	Set(index int, val any) error
	Set2(key string, val any) error

	Value(index int) any
	Value2(key string) any

	String(index int) string
	Int64(index int) int64
	Int(index int) int
	Int8(index int) int8
	Int16(index int) int16
	Int32(index int) int32
	Uint64(index int) uint64
	Uint(index int) uint
	Uint8(index int) uint8
	Uint16(index int) uint16
	Uint32(index int) uint32
	Float64(index int) float64
	Float32(index int) float32
	Bool(index int) bool

	ValueOrDefault(index int, d any) any
	ValueOrDefault2(key string, d any) any

	StringOrDefault(index int, d string) string
	Int64OrDefault(index int, d int64) int64
	IntOrDefault(index int, d int) int
	Int8OrDefault(index int, d int8) int8
	Int16OrDefault(index int, d int16) int16
	Int32OrDefault(index int, d int32) int32
	Uint64OrDefault(index int, d uint64) uint64
	UintOrDefault(index int, d uint) uint
	Uint8OrDefault(index int, d uint8) uint8
	Uint16OrDefault(index int, d uint16) uint16
	Uint32OrDefault(index int, d uint32) uint32
	Float64OrDefault(index int, d float64) float64
	Float32OrDefault(index int, d float32) float32
	BoolOrDefault(index int, d bool) bool

	TryValue(index int) (val any, ok bool)
	TryValue2(key string) (val any, ok bool)

	TryString(index int) (res string, ok bool)
	TryInt64(index int) (int64, bool)
	TryInt(index int) (int, bool)
	TryInt8(index int) (int8, bool)
	TryInt16(index int) (int16, bool)
	TryInt32(index int) (int32, bool)
	TryUint64(index int) (uint64, bool)
	TryUint(index int) (uint, bool)
	TryUint8(index int) (uint8, bool)
	TryUint16(index int) (uint16, bool)
	TryUint32(index int) (uint32, bool)
	TryFloat64(index int) (float64, bool)
	TryFloat32(index int) (float32, bool)
	TryBool(index int) (bool, bool)
}
