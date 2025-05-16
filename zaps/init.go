package zaps

import (
	"github.com/gomsr/atom-msr/kg"
	zapx "github.com/gomsr/atom-zap"
	"go.uber.org/zap"
)

func InitZap() *zap.Logger {
	kg.L = zapx.Zap(kg.C.Zap)
	zap.ReplaceGlobals(kg.L)
	return kg.L
}
