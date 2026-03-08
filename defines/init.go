package defines

import (
	"context"
	"github.com/skypbc/public/defines/dconfig"
)

func InitConfigValues(ctx context.Context) (err error) {
	callbacks := []func(ctx context.Context) error{
		dconfig.Init,
	}

	for _, fn := range callbacks {
		if err := fn(ctx); err != nil {
			return err
		}
	}

	return nil
}
