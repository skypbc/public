package engine

import (
	"context"

	"github.com/skypbc/public/defines/dobjects"
	"github.com/skypbc/public/include"
)

var NewEngine func(objectManager IObjectManager) IEngine
var SetEngine func(ctx context.Context, engine IEngine) context.Context
var ExtractEngine func(context.Context) include.IEngine

var Init func(ctx context.Context, engine IEngine) (context.Context, error) = initEngine

func initEngine(ctx context.Context, engine IEngine) (context.Context, error) {
	ctx = SetEngine(ctx, engine)
	err := dobjects.Init(ctx)
	if err != nil {
		return ctx, err
	}
	return ctx, nil
}
