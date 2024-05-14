package settings_api

import (
	"gin_blog/global"
	response "gin_blog/models/common"
	"github.com/gin-gonic/gin"
)

type SettingsUri struct {
	Name string `uri:"name"`
}

// 显示某一项的配置信息
func (SettingsApi) SettingsInfoView(c *gin.Context) {
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}
	//todo 后续增加对应的配置选项
	switch cr.Name {
	case "email":
		response.OkWithData(global.Config.QiNiu, c)
	case "qq":
		response.OkWithData(global.Config.QiNiu, c)
	case "qiniu":
		response.OkWithData(global.Config.QiNiu, c)
	case "jwt":
		response.OkWithData(global.Config.QiNiu, c)
	default:
		response.FailWithMessage("没有对应的配置信息", c)
	}
}
