package engine

import (
	"context"
	"github.com/skypbc/public/include"
)

var NewEngine func(objectManager IObjectManager) IEngine
var SetEngine func(ctx context.Context, engine IEngine) context.Context
var ExtractEngine func(context.Context) include.IEngine
