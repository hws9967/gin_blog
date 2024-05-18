package router

import (
	"gin_blog/api"
	"github.com/gin-gonic/gin"
)

func (ApiRouter) AdvertRouter(Router *gin.RouterGroup) {
	userRouterApi := api.ApiGroupApp.UserApi
	{
		Router.GET("users", userRouterApi.UserListView)
	}
}
