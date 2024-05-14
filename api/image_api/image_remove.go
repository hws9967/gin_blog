package image_api

import (
	"fmt"
	"gin_blog/global"
	"gin_blog/models"
	response "gin_blog/models/common"
	"github.com/gin-gonic/gin"
)

type ImageBaseRequest struct {
	ID uint `uri:"id"`
}

func (ImageApi) ImageRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	var imageList []models.BannerModel
	count := global.DB.Find(&imageList, cr.IDList).RowsAffected
	if count == 0 {
		response.FailWithMessage("文件不存在", c)
		return
	}
	global.DB.Delete(&imageList)
	response.OkWithMessage(fmt.Sprintf("共删除 %d 张图片", count), c)

}
