package core

import (
	"fmt"
	"gin_blog/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

func Gorm() *gorm.DB {
	return MysqlConnect()
}

func MysqlConnect() *gorm.DB {
	if global.CONFIG.Mysql.Host == "" {
		log.Println("未配置mysql，取消gorm链接")
		fmt.Println("未配置mysql，取消gorm链接")
		return nil
	}

	dsn := global.CONFIG.Mysql.Dsn()

	var mysqlLogger logger.Interface
	if global.CONFIG.System.Env == "dev" {
		// 开发环境显示所有的sql
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error) // 只打印错误的sql
	}
	global.MysqlLog = logger.Default.LogMode(logger.Info)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   mysqlLogger, //配置日志级别，打印出所有的sql
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		global.LOG.Fatalf("mysql连接失败 %s", err.Error())
	}
	sqlDB, _ := db.DB()

	sqlDB.SetMaxIdleConns(10)           // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)          // 最多可容纳
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接最大复用时间
	return db
}
