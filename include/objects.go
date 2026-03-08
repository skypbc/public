package include

type ObjectManagerItem struct {
	Key    string
	Index  int
	Object any
}

type IObjectManager interface {
	Index(key string, create ...bool) int
	Key(index int) string

	SetByKey(key string, obj any) error
	SetByIndex(index int, obj any) error

	GetByKey(key string) any
	GetByIndex(index int) any

	List() []ObjectManagerItem
}
