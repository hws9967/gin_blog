package settings_api

import (
	"gin_blog/config"
	"gin_blog/core"
	"gin_blog/global"
	response "gin_blog/models/common"
	"github.com/gin-gonic/gin"
)

// todo
func (SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}
	switch cr.Name {
	case "email":
		var info config.Email
		err = c.ShouldBindJSON(&info)
		if err != nil {
			response.FailWithCode(response.ArgumentError, c)
			return
		}
		global.Config.Email = info
	//case "qq":
	//var info config.QQ
	//err = c.ShouldBindJSON(&info)
	//if err != nil {
	//	response.FailWithCode(response.ArgumentError, c)
	//	return
	//}
	//	global.Config.QQ = info
	case "qiniu":
		var info config.QiNiu
		err = c.ShouldBindJSON(&info)
		if err != nil {
			response.FailWithCode(response.ArgumentError, c)
			return
		}
		global.Config.QiNiu = info
	case "jwt":
		var info config.Jwy
		err = c.ShouldBindJSON(&info)
		if err != nil {
			response.FailWithCode(response.ArgumentError, c)
			return
		}
		global.Config.Jwy = info
	default:
		response.FailWithMessage("没有对应的配置信息", c)
		return
	}

	core.SetYaml()
	response.OkWith(c)
}
