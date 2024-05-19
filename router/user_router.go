package router

import (
	"gin_blog/api"
	"gin_blog/middleware"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var store = cookie.NewStore([]byte("HyvCD89g3VDJ9646BFGEh37GFJ"))

func (ApiRouter) UserRouter(Router *gin.RouterGroup) {
	userRouterApi := api.ApiGroupApp.UserApi
	{
		Router.POST("email_login", userRouterApi.EmailLoginView)
		Router.GET("users", middleware.JwtAuth(), userRouterApi.UserListView)
		Router.POST("user_bind_email", userRouterApi.UserBindEmailView)
		Router.PUT("user_role", middleware.JwtAdmin(), userRouterApi.UserUpdateRoleView)
		Router.POST("logout", middleware.JwtAuth(), userRouterApi.LogoutView)
	}
}
