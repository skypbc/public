package objects

import (
	"github.com/skypbc/public/include"
)

type IObjectManager = include.IObjectManager

var NewObjectManager func() IObjectManager
