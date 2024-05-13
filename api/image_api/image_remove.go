package image_api

import (
	"fmt"
	"gin_blog/global"
	"gin_blog/models"
	response "gin_blog/models/common"
	"gin_blog/models/image_type"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ImageBaseRequest struct {
	ID uint `uri:"id"`
}

func (ImageApi) ImageRemoveView(c *gin.Context) {
	var cr ImageBaseRequest
	c.ShouldBindUri(&cr)
	//判断当前图片是否有人使用，使用就必须提示用户
	var imageModel models.ImageModel
	err := global.DB.Take(&imageModel, cr.ID).Error
	if err != gorm.ErrRecordNotFound {
		response.FailWithMessage("删除失败，图片不存在", c)
		return
	}
	var count int
	switch imageModel.ImageType {
	case image_type.TypeCover:
		//图片与封面有关
		global.DB.Where(&imageModel).Preload("ArticleModels").Find(&imageModel)
		count = len(imageModel.ArticleModels)
		if count > 0 {
			response.FailWithMessage(fmt.Sprintf("图片与文章封面有关联%s篇", count), c)
			return
		}
	case image_type.TypeAvatar:
		//图片与头像有关
		global.DB.Where(&imageModel).Preload("AuthModels").Find(&imageModel)
		count = len(imageModel.AuthModels)
		if count > 0 {
			response.FailWithMessage(fmt.Sprintf("图片与头像有关联%s篇", count), c)
			return
		}
	}

	err = global.DB.Delete(&imageModel).Error
	if err != nil {
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
