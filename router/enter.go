package router

import (
	"gin_blog/api"
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

	settingsApi := api.ApiGroupApp.SettingsApi
	// 系统设置
	router.GET("infoview", settingsApi.InfoView)

	routerGroupApp := RouterGroup{router}
	routerGroupApp.SettingsRouter()
	return router
}
