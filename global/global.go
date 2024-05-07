package global

import (
	"gin_blog/config"
	"github.com/cc14514/go-geoip2"
	"github.com/go-redis/redis/v8"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
)

var (
	Config   *config.Config
	DB       *gorm.DB
	DBList   map[string]*gorm.DB
	REDIS    *redis.Client
	CONFIG   config.Server
	LOG      *logrus.Logger
	ESClient *elastic.Client
	MysqlLog logger.Interface
	lock     sync.RWMutex
	AddrDB   *geoip2.DBReader
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}
