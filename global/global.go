package global

import (
	"gin_blog/config"
	"github.com/cc14514/go-geoip2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Config   *config.Config
	DB       *gorm.DB
	Log      *logrus.Logger
	MysqlLog logger.Interface
	AddrDB   *geoip2.DBReader
)
