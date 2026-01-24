package cmds

import (
	"context"
	"github.com/skypbc/public/defines/derrs"
)

// BaseCmd является основой для описания своих собственных комманд.
//
// Важно: Команда не может и не должна хранить какие-либо данные. Так как один экземпляр используется во
// множестве горутин... Все методы следует считать статическими. А экземпляры комманд синглетонами.
type BaseCmd struct{}

// Name возвращает имя команды
func (b *BaseCmd) Name() string {
	panic(derrs.NewNotImplementedError())
}

// Kind возвращает вид команды
func (b *BaseCmd) Kind() Kind {
	panic(derrs.NewNotImplementedError())
}

// Category возвращает категорию команды
func (b *BaseCmd) Category() []Category {
	panic(derrs.NewNotImplementedError())
}

// Action возвращает действие команды
func (b *BaseCmd) Action() Action {
	panic(derrs.NewNotImplementedError())
}

// CheckAccess проверяет права на выполнение команды для заданного ивента
func (b *BaseCmd) CheckAccess(ctx context.Context, event IEvent) (ok bool, err error) {
	return false, derrs.NewNotImplementedError()
}

// Execute выполняет команду для заданного ивента
func (b *BaseCmd) Execute(ctx context.Context, event IEvent) (IEvent, error) {
	return nil, derrs.NewNotImplementedError()
}

func (b *BaseCmd) CheckOnResultAccess(ctx context.Context, event IEvent) (ok bool, err error) {
	return false, derrs.NewNotImplementedError()
}

func (b *BaseCmd) OnResult(ctx context.Context, event IEvent) (IEvent, error) {
	return nil, derrs.NewNotImplementedError()
}

func (b *BaseCmd) OnResultError(ctx context.Context, event IEvent) (IEvent, error) {
	return nil, derrs.NewNotImplementedError()
}
