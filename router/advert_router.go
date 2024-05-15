package router

import (
	"gin_blog/api"
	"github.com/gin-gonic/gin"
)

func (ApiRouter) AdvertRouter(Router *gin.RouterGroup) {
	advertRouterApi := api.ApiGroupApp.AdvertApi
	{
		Router.GET("adverts", advertRouterApi.AdvertListView)
		Router.POST("adverts", advertRouterApi.AdvertCreateView)
		Router.PUT("adverts/:id", advertRouterApi.AdvertUpdateView)
		Router.DELETE("adverts", advertRouterApi.AdvertRemoveView)
	}
}
