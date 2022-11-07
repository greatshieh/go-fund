package global

import (
	"sync"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"gospider/configs"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GPA_DB                  *gorm.DB
	GPA_DBList              map[string]*gorm.DB
	GPA_CONFIG              configs.Server
	GPA_VP                  *viper.Viper
	GPA_LOG                 *zap.Logger
	GPA_Concurrency_Control = &singleflight.Group{}
	lock                    sync.RWMutex
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return GPA_DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := GPA_DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}
