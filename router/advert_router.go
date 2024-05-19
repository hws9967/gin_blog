package router

import (
	"gin_blog/api"
	"github.com/gin-gonic/gin"
)

func (ApiRouter) AdvertRouter(Router *gin.RouterGroup) {
	advertRouterApi := api.ApiGroupApp.AdvertApi
	{
		Router.POST("adverts", advertRouterApi.AdvertCreateView)
		Router.GET("adverts", advertRouterApi.AdvertListView)
		Router.PUT("adverts/:id", advertRouterApi.AdvertUpdateView)
		Router.DELETE("adverts", advertRouterApi.AdvertRemoveView)
	}
}
