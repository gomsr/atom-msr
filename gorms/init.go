package gorms

import (
	gormx "github.com/gomsr/atom-gorm"
	"github.com/gomsr/atom-gorm/gconfig"
	"github.com/gomsr/atom-gorm/initialize/gormc"
	"gorm.io/gorm"
	"log"
	"os"
)

func InitDB(c gconfig.DbServer) *gorm.DB {
	gormc.Gorm.SetLogger(NewWriter(log.New(os.Stdout, "\r\n", log.LstdFlags)))
	return gormx.InitDB(c)
}
