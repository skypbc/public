package include

type EngineCtxKey struct{}

type IEngine interface {
	Config() IEngineConfig
	SetConfig(config IEngineConfig, force bool) error

	ObjectManager() IObjectManager
}
