package atommsr

import (
	"github.com/gomsr/atom-msr/gorms"
	"github.com/gomsr/atom-msr/kg"
	"github.com/gomsr/atom-msr/zaps"
	viperx "github.com/gomsr/atom-viper"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

func InitV0(conf any, path ...string) {
	viperx.InitViperWithConf(&kg.C, conf, path...)
	kg.L = zaps.InitZap()
	kg.DB = gorms.InitDB(kg.C.DbServer)
}

// Init @Deprecated
func Init(path ...string) (*viper.Viper, *zap.Logger, *gorm.DB) {
	kg.V = viperx.InitViper(&kg.C, path...)
	return coreInit()
}

func InitV2(defaultPath ...string) (*viper.Viper, *zap.Logger, *gorm.DB) {
	configFile := detConfig(defaultPath)

	kg.V = viperx.InitViper(&kg.C, configFile)
	return coreInit()
}

func InitWithConf(conf any, path ...string) (*viper.Viper, *zap.Logger, *gorm.DB) {
	kg.V = viperx.InitViperWithConf(&kg.C, conf, path...)
	return coreInit()
}

func InitWithConfV2(conf any, defaultPath ...string) (*viper.Viper, *zap.Logger, *gorm.DB) {
	configFile := detConfig(defaultPath)

	kg.V = viperx.InitViperWithConf(&kg.C, conf, configFile)
	return coreInit()
}

func coreInit() (*viper.Viper, *zap.Logger, *gorm.DB) {
	kg.L = zaps.InitZap()
	kg.DB = gorms.InitDB(kg.C.DbServer)
	return kg.V, kg.L, kg.DB
}

func detConfig(defaultPath []string) string {
	con := os.Getenv("config")
	if len(con) != 0 {
		return con
	}

	if len(defaultPath) == 0 {
		panic("config file not found")
	}
	return defaultPath[0]
}
