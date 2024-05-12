package router

import (
	"gin_blog/api"
	"github.com/gin-gonic/gin"
)

func (ApiRouter) InitImage(Router *gin.RouterGroup) {
	imageRouterApi := api.ApiGroupApp.ImageApi
	{
		Router.POST("images", imageRouterApi.ImageUploadView)
		Router.GET("images", imageRouterApi.ImageListView) // 图片列表
	}
}
