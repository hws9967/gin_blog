package router

import (
	"gin_blog/api"
	"github.com/gin-gonic/gin"
)

func (ApiRouter) InitSettings(Router *gin.RouterGroup) {
	routerApi := api.ApiGroupApp.SettingsApi
	{
		Router.GET("site_info", routerApi.SettingsInfoView)
	}
}

func (router RouterGroup) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	router.GET("", settingsApi.SettingsInfoView)
}

func (ApiRouter) SettingsRouter1(Router *gin.RouterGroup) {
	routerApi := api.ApiGroupApp.SettingsApi
	{
		Router.GET("/site", routerApi.InfoView)
	}
}
