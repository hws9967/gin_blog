package settings_api

import (
	response "gin_blog/models/common"
	"github.com/gin-gonic/gin"
)

func (SettingsApi) InfoView(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "settings info this is infoview"})
	//return
}

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "settings info"})
}

func (SettingsApi) InfoOk(c *gin.Context) {
	response.Ok(c)
}

func (SettingsApi) OkDetailed(c *gin.Context) {
	response.OkWithDetailed(map[string]string{
		"name":    "张三",
		"message": "ok",
	}, "信息获取成功", c)
}
