package main

import (
	"gin_blog/core"
	_ "gin_blog/docs"
	"gin_blog/global"
	"gin_blog/router"
)

// @title gvb_server API文档
// @version 1.0
// @description gvb_server API文档
// @host 127.0.0.1:8080
// @BasePath /
func main() {

	// 读取配置文件
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()
	// 连接数据库
	global.DB = core.InitGorm()
	// 初始化路由
	router := router.InitRouter()
	// 连接redis
	global.Redis = core.ConnectRedis()
	// 启动服务器

	addr := global.Config.System.Addr()
	err := router.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
}
