package main

import (
	"gin_blog/core"
	"gin_blog/global"
	routers "gin_blog/router"
)

func main() {
	global.CONFIG = core.ConfInit() //读取配置文件
	global.LOG = core.InitLogger()  //初始化日志
	/*	global.LOG.Warnln("这是警告")
		global.LOG.Infoln("这是信息")
		global.LOG.Errorln("这是错误")*/
	global.DB = core.Gorm() // 初始化数据库
	router := routers.InitRouter()
	addr := global.CONFIG.System.GetAddr()
	global.LOG.Infof("服务器启动成功:端口%s", addr)
	router.Run(":8080")
}
