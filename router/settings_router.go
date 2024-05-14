package router

import (
	"gin_blog/api"
	"github.com/gin-gonic/gin"
)

func (ApiRouter) InitSettings(Router *gin.RouterGroup) {
	routerApi := api.ApiGroupApp.SettingsApi
	{
		Router.GET("setting/site", routerApi.SettingsSiteInfoView)
		Router.PUT("setting/site", routerApi.SettingsSiteUpdateView)
		Router.GET("settings/:name", routerApi.SettingsInfoView)
		Router.PUT("settings/:name", routerApi.SettingsInfoUpdateView)
	}
}
