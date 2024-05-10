package router

import (
	"gin_blog/global"
	"github.com/gin-gonic/gin"
)

type ApiRouter struct {
}

var apiRouterGroup = new(ApiRouter)

type RouterGroup struct {
	*gin.Engine
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.CONFIG.System.Env)
	router := gin.Default()

	// 系统设置
	PublicGroup := router.Group("api")
	apiRouterGroup.InitSettings(PublicGroup)
	apiRouterGroup.InitImage(PublicGroup)
	return router
}
