package router

import (
	"gin_blog/api"
	"github.com/gin-gonic/gin"
)

func (ApiRouter) InitImage(Router *gin.RouterGroup) {
	imageRouterApi := api.ApiGroupApp.ImageApi
	{
		Router.POST("images", imageRouterApi.ImageUploadView)    // 上传图片
		Router.POST("image", imageRouterApi.ImageUploadDataView) // 上传图片至七牛云
		Router.GET("images", imageRouterApi.ImageListView)       // 图片列表
		Router.GET("image_names", imageRouterApi.ImageNameListView)
		Router.DELETE("images", imageRouterApi.ImageRemoveView) // 删除图片
		Router.PUT("images", imageRouterApi.ImageUpdateView)    // 删除图片// 图片列表
	}
}
