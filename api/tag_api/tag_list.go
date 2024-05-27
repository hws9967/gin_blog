package tag_api

import (
	"gin_blog/models"
	response "gin_blog/models/common"
	"gin_blog/services/common"
	"github.com/gin-gonic/gin"
)

func (TagApi) TagListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	list, count, _ := common.ComList(models.TagModel{}, common.Option{
		PageInfo: cr,
	})

	// 返回结果
	response.OkWithList(list, count, c)
}
