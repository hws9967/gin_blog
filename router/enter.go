package router

import (
	"gin_blog/api"
	"gin_blog/global"
	"github.com/gin-gonic/gin"
)

type ApiRouter struct {
}

var apiRouterGroup = new(ApiRouter)

func InitRouter() *gin.Engine {
	gin.SetMode(global.CONFIG.System.Env)
	router := gin.Default()

	settingsApi := api.ApiGroupApp.SettingsApi
	// 系统设置
	router.GET("", settingsApi.InfoView)

	return router
}
