package router

import (
	"gin_blog/api"
	"gin_blog/middleware"
	"github.com/gin-gonic/gin"
)

func (ApiRouter) MessageRouter(Router *gin.RouterGroup) {
	messageRouterApi := api.ApiGroupApp.MessageApi
	{
		Router.POST("messages", middleware.JwtAuth(), messageRouterApi.MessageCreateView)
		Router.GET("messages_all", messageRouterApi.MessageListAllView)
		Router.GET("messages", middleware.JwtAuth(), messageRouterApi.MessageListView)
		Router.POST("messages_record", middleware.JwtAuth(), messageRouterApi.MessageRecordView)
		Router.GET("message_users", middleware.JwtAdmin(), messageRouterApi.MessageUserListView)
		Router.DELETE("message_users", middleware.JwtAuth(), messageRouterApi.MessageRecordRemoveView)
		Router.GET("message_users/user", middleware.JwtAdmin(), messageRouterApi.MessageUserListByUserView)
		Router.GET("message_users/record", middleware.JwtAdmin(), messageRouterApi.MessageUserRecordView)
		Router.GET("message_users/me", middleware.JwtAuth(), messageRouterApi.MessageUserListByMeView)
		Router.GET("message_users/record/me", middleware.JwtAuth(), messageRouterApi.MessageUserRecordByMeView)

	}
}
