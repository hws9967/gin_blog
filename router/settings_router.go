package router

import (
	"gin_blog/api"
	"github.com/gin-gonic/gin"
)

func (ApiRouter) InitSettings(Router *gin.RouterGroup) {
	routerApi := api.ApiGroupApp.SettingsApi
	{
		Router.GET("site_info", routerApi.InfoView)
		Router.GET("ok", routerApi.InfoOk)
		Router.GET("ok_with_detailed", routerApi.OkDetailed)
	}
}
