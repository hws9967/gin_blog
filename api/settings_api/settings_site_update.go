package settings_api

import (
	"gin_blog/config"
	"gin_blog/core"
	"gin_blog/global"
	response "gin_blog/models/common"
	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsSiteUpdateView(c *gin.Context) {
	var info config.SiteInfo
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	global.Config.SiteInfo = info
	core.SetYaml()
	response.OkWithMessage("修改成功", c)
}
