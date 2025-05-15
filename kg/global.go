package kg

import (
	"github.com/gomsr/atom-msr/vconfig"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	C  vconfig.Server
	V  *viper.Viper
	L  *zap.Logger
	DB *gorm.DB
)
