package gorms

import (
	"fmt"
	"github.com/gomsr/atom-gorm/initialize/gormc"
	"github.com/gomsr/atom-msr/kg"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm/logger"
)

type writer struct {
	logger.Writer
}

// NewWriter writer 构造函数
func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

// Printf 格式化打印日志
func (w *writer) Printf(message string, data ...interface{}) {
	var logType string
	switch kg.C.DbServer.DbType {
	case "mysql":
		logType = kg.C.Mysql.LogType
	case "pgsql":
		logType = kg.C.Pgsql.LogType
	}

	if logType == gormc.Zap {
		kg.L.Info(fmt.Sprintf(message+"\n", data...))
	} else if logType == gormc.GoZero {
		logx.Debug(fmt.Sprintf(message+"\n", data...))
	}

	w.Writer.Printf(message, data...)
}
