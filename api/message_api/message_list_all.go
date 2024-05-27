package message_api

import (
	"gin_blog/models"
	response "gin_blog/models/common"
	"gin_blog/services/common"
	"github.com/gin-gonic/gin"
)

// MessageListAllView 消息列表
// @Tags 消息管理
// @Summary 消息列表
// @Description 消息列表
// @Router /api/messages_all [get]
// @Param token header string  true  "token"
// @Param data query models.PageInfo    false  "查询参数"
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.MessageModel]}
func (MessageApi) MessageListAllView(c *gin.Context) {
	var cr models.PageInfo
	if err := c.ShouldBindQuery(&cr); err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}
	list, count, _ := common.ComList(models.MessageModel{}, common.Option{
		PageInfo: cr,
	})
	response.OkWithList(list, count, c)
}
