package router

import (
	"gin_blog/global"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

type ApiRouter struct {
}

var apiRouterGroup = new(ApiRouter)

type RouterGroup struct {
	*gin.Engine
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	//配置swag
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	// 系统设置
	PublicGroup := router.Group("api")
	apiRouterGroup.InitSettings(PublicGroup)
	apiRouterGroup.InitImage(PublicGroup)
	apiRouterGroup.AdvertRouter(PublicGroup)
	apiRouterGroup.MenuRouter(PublicGroup)
	return router
}
