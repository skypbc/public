package defines

import (
	"context"
	"github.com/skypbc/public/defines/dconfig"
	"github.com/skypbc/public/defines/dobjects"
)

func InitConfigValues(ctx context.Context) (err error) {
	callbacks := []func(ctx context.Context) error{
		dconfig.Init,
		dobjects.Init,
	}

	for _, fn := range callbacks {
		if err := fn(ctx); err != nil {
			return err
		}
	}

	return nil
}
