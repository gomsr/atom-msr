package atommsr

import (
	gormx "github.com/gomsr/atom-gorm"
	"github.com/gomsr/atom-msr/kg"
	viperx "github.com/gomsr/atom-viper"
	zapx "github.com/gomsr/atom-zap"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func InitZap() *zap.Logger {
	kg.L = zapx.Zap(kg.C.Zap)
	zap.ReplaceGlobals(kg.L)
	return kg.L
}

func InitV0(conf any, path ...string) {
	viperx.InitViperWithConf(&kg.C, conf, path...)
	kg.L = InitZap()
	kg.DB = gormx.InitDB(kg.C.DbServer)
}

func Init(path ...string) (*viper.Viper, *zap.Logger, *gorm.DB) {
	kg.V = viperx.InitViper(&kg.C, path...)
	kg.L = InitZap()
	kg.DB = gormx.InitDB(kg.C.DbServer)

	return kg.V, kg.L, kg.DB
}

func InitWithConf(conf any, path ...string) (*viper.Viper, *zap.Logger, *gorm.DB) {
	kg.V = viperx.InitViperWithConf(&kg.C, conf, path...)
	kg.L = InitZap()
	kg.DB = gormx.InitDB(kg.C.DbServer)

	return kg.V, kg.L, kg.DB
}
