package settings_api

import "github.com/gin-gonic/gin"

func (SettingsApi) InfoView(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "settings info this is infoview"})
	//return
}

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "settings info"})
}
