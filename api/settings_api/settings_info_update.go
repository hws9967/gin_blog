package settings_api

import (
	"fmt"
	"gin_blog/config"
	"gin_blog/core"
	"gin_blog/global"
	response "gin_blog/models/common"
	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	var cr config.SiteInfo
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithMessage("传入类型错误", c)
		return
	}
	fmt.Printf("%#v\n", cr)

	global.CONFIG.SiteInfo = cr
	core.SetYaml(global.CONFIG)

	response.OkWithMessage("修改成功", c)
	return
}
