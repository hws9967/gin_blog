package router

import (
	"gin_blog/api"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var store = cookie.NewStore([]byte("HyvCD89g3VDJ9646BFGEh37GFJ"))

func (ApiRouter) UserRouter(Router *gin.RouterGroup) {
	advertRouterApi := api.ApiGroupApp.AdvertApi
	{
		Router.GET("adverts", advertRouterApi.AdvertListView)
		Router.POST("adverts", advertRouterApi.AdvertCreateView)
		Router.PUT("adverts/:id", advertRouterApi.AdvertUpdateView)
		Router.DELETE("adverts", advertRouterApi.AdvertRemoveView)
	}
}
