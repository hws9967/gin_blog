package router

import (
	"gin_blog/api"
	"github.com/gin-gonic/gin"
)

func (ApiRouter) InitSettings(Router *gin.RouterGroup) {
	routerApi := api.ApiGroupApp.SettingsApi
	{
		Router.GET("site_info", routerApi.InfoView)
	}
}
