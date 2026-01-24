package include

type IObjectManager interface {
	Get(key int) any
	Set(key int, obj any, force bool) error
}
