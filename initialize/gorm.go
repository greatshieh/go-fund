package initialize

import (
	"os"

	FundModel "gospider/app/fund/model"
	"gospider/global"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Gorm 初始化数据库并产生数据库全局变量
func Gorm() *gorm.DB {
	switch global.GPA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	default:
		return GormMysql()
	}
}

// RegisterTables 注册数据库表专用
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		FundModel.FundModel{},
		FundModel.HoldInfo{},
		FundModel.YearlyGainsModel{},
		// ZhihuModel.ZhihuModel{},
	)
	if err != nil {
		global.GPA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GPA_LOG.Info("register table success")
}
