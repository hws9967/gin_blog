package main

import (
	"gin_blog/core"
	"gin_blog/global"
	"gin_blog/router"
)

func main() {
	//global.CONFIG = core.ConfInit() //读取配置文件
	//global.LOG = core.InitLogger()  //初始化日志
	///*	global.LOG.Warnln("这是警告")
	//	global.LOG.Infoln("这是信息")
	//	global.LOG.Errorln("这是错误")*/
	//global.DB = core.Gorm() // 初始化数据库
	//router := routers.InitRouter()
	//addr := global.CONFIG.System.GetAddr()
	//global.LOG.Infof("服务器启动成功:端口%s", addr)
	//router.Run(":8080")

	// 读取配置文件
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()
	// 连接数据库
	global.DB = core.InitGorm()
	// 初始化路由
	router := router.InitRouter()
	// 启动服务器

	addr := global.Config.System.Addr()
	err := router.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
}
