package image_api

import (
	"gin_blog/models"
	response "gin_blog/models/common"
	"gin_blog/models/image_type"
	"gin_blog/services/commont_ser"
	"github.com/gin-gonic/gin"
)

type ImageIDResponse struct {
	Path string `json:"path"` // 图片路径
	ID   uint   `json:"id"`
	Name string `json:"name"` // 图片名称
}
type ImageIDRequest struct {
	Type int64 `form:"type"`
}

func (ImageApi) ImageIdListView(c *gin.Context) {
	var cr ImageIDRequest
	c.ShouldBindQuery(&cr)
	if cr.Type == 0 {
		cr.Type = 1
	}
	list, _, err := commont_ser.CommonList(&models.ImageModel{
		ImageType: image_type.ImageType(cr.Type),
	}, commont_ser.ListOption{
		Where: []string{
			"not_use = false",
		},
	})
	if err != nil {
		response.FailWithMessage("意外错误", c)
		return
	}
	var imageList []ImageIDResponse
	for _, model := range list {
		imageList = append(imageList, ImageIDResponse{
			Path: model.Path,
			ID:   model.ID,
			Name: model.Name,
		})
	}
	response.OkWithDetailed(imageList, "成功", c)
}
