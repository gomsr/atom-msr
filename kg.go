package atommsr

import (
	"github.com/gomsr/atom-msr/gorms"
	"github.com/gomsr/atom-msr/kg"
	"github.com/gomsr/atom-msr/zaps"
	viperx "github.com/gomsr/atom-viper"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func InitV0(conf any, path ...string) {
	viperx.InitViperWithConf(&kg.C, conf, path...)
	kg.L = zaps.InitZap()
	kg.DB = gorms.InitDB(kg.C.DbServer)
}

func Init(path ...string) (*viper.Viper, *zap.Logger, *gorm.DB) {
	kg.V = viperx.InitViper(&kg.C, path...)
	kg.L = zaps.InitZap()
	kg.DB = gorms.InitDB(kg.C.DbServer)

	return kg.V, kg.L, kg.DB
}

func InitWithConf(conf any, path ...string) (*viper.Viper, *zap.Logger, *gorm.DB) {
	kg.V = viperx.InitViperWithConf(&kg.C, conf, path...)
	kg.L = zaps.InitZap()
	kg.DB = gorms.InitDB(kg.C.DbServer)
	return kg.V, kg.L, kg.DB
}
