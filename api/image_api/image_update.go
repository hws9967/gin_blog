package image_api

import (
	"gin_blog/global"
	"gin_blog/models"
	response "gin_blog/models/common"
	"gin_blog/plugins/logstash/lm"
	"github.com/gin-gonic/gin"
)

type ImageUpdateRequest struct {
	Sort int    `json:"sort"`
	Name string `json:"name"`
}

func (ImageApi) ImageUpdate(c *gin.Context) {
	id := c.Param("id")
	var cr ImageUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithMessage("类型错误", c)
		return
	}
	var image models.ImageModel
	// 名字不能重名
	err = global.DB.Take(&image, "name = ?", cr.Name).Error
	if err == nil {
		// 找到了
		response.FailWithMessage("图片名称不能重复", c)
		return
	}
	lm.LoggerApp.Send("更新图片名称", c)
	global.DB.Take(&image, id).UpdateColumns(map[string]any{
		"sort": cr.Sort,
		"name": cr.Name,
	})
	response.OkWithMessage("修改成功", c)
	return

}
