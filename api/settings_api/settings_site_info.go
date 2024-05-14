package settings_api

import (
	"gin_blog/global"
	response "gin_blog/models/common"
	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsSiteInfoView(c *gin.Context) {
	response.OkWithData(global.Config.SiteInfo, c)
}
