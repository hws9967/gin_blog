package settings_api

import (
	"gin_blog/global"
	response "gin_blog/models/common"
	"github.com/gin-gonic/gin"
)

// 读取yaml中的配置文件
func (SettingsApi) SettingsInfoView(c *gin.Context) {
	response.OkWithData(global.CONFIG.SiteInfo, c)
}
