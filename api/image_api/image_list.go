package image_api

import (
	"gin_blog/models"
	response "gin_blog/models/common"
	"gin_blog/models/image_type"
	"gin_blog/services/commont_ser"
	"github.com/gin-gonic/gin"
)

type ImageListRequest struct {
	models.PageInfo
	ImageType int `json:"image_type" form:"image_type"`
}

// ImageListView 文件列表
func (i ImageApi) ImageListView(c *gin.Context) {

	var cr ImageListRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		response.FailWithMessage("错误", c)
		return
	}

	list, count, err := commont_ser.CommonList(&models.ImageModel{
		ImageType: image_type.ImageType(cr.ImageType),
	}, commont_ser.ListOption{
		Page: cr.PageInfo,
		Where: []string{
			"not_use = false",
		},
	})
	if err != nil {
		response.FailWithMessage("意外错误", c)
		return
	}
	response.OkWithDetailed(commont_ser.CommonResponse{List: list, Count: count}, "成功", c)
}
