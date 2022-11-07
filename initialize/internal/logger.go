package internal

import (
	"fmt"

	"gospider/global"

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
	var logZap bool
	switch global.GPA_CONFIG.System.DbType {
	case "mysql":
		logZap = global.GPA_CONFIG.MySQL.LogZap
	case "pgsql":
		logZap = global.GPA_CONFIG.PgSQL.LogZap
	}
	if logZap {
		global.GPA_LOG.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}
