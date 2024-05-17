package router

import (
	"gin_blog/api"
	"github.com/gin-gonic/gin"
)

func (ApiRouter) MenuRouter(Router *gin.RouterGroup) {
	advertRouterApi := api.ApiGroupApp.MenuAPi
	{
		Router.GET("menus", advertRouterApi.MenuListView)
		Router.POST("menus", advertRouterApi.MenuCreateView)
		Router.GET("menus_names", advertRouterApi.MenuNameList)
		Router.PUT("menus/:id", advertRouterApi.MenuUpdateView)
		Router.GET("menus/:id", advertRouterApi.MenuDetailView)
		Router.DELETE("menus", advertRouterApi.MenuRemoveView)
		Router.GET("menus/detail", advertRouterApi.MenuDetailByPathView)
	}
}
