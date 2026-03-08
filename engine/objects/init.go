package objects

import (
	"context"

	"github.com/skypbc/public/include"
)

type IObjectManager = include.IObjectManager

var NewObjectManager func() IObjectManager
var ExtractObjectManager func(ctx context.Context) IObjectManager
