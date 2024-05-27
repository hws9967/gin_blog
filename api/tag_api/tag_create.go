package tag_api

import (
	"gin_blog/global"
	"gin_blog/models"
	response "gin_blog/models/common"
	"github.com/gin-gonic/gin"
)

type TagRequest struct {
	Title string `json:"title" binding:"require" msg:"请输入标题" structs:"title"`
}

// TagCreateView 发布标签
// @Tags 标签管理
// @Summary 发布标签
// @Description 发布标签
// @Param data body TagRequest  true  "查询参数"
// @Param token header string  true  "token"
// @Router /api/tags [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (TagApi) TagCreateView(c *gin.Context) {
	var cr TagRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}

	// 业务逻辑 重复的判断
	var tag models.TagModel
	err = global.DB.Take(&tag, "title = ?", cr.Title).Error

	if err == nil {
		response.FailWithMessage("标签已存在", c)
		return
	}
	// 创建标签
	err = global.DB.Create(&models.TagModel{
		Title: cr.Title,
	}).Error

	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("创建标签失败", c)
		return
	}

	response.OkWithMessage("创建标签成功", c)
}
