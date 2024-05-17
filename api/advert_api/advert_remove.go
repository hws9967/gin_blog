package advert_api

import (
	"fmt"
	"gin_blog/global"
	"gin_blog/models"
	response "gin_blog/models/common"
	"github.com/gin-gonic/gin"
)

// AdvertRemoveView 批量删除广告
// @Tags 广告管理
// @Summary 批量删除广告
// @Description 批量删除广告
// @Param data body models.RemoveRequest    true  "广告id列表"
// @Param token header string  true  "token"
// @Router /api/adverts [delete]
// @Produce json
// @Success 200 {object} res.Response{}
func (AdvertApi) AdvertRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	var advertList []models.AdvertModel
	count := global.DB.Find(&advertList, cr.IDList).RowsAffected

	if count == 0 {
		response.FailWithMessage("删除失败,广告不存在", c)
		return
	}

	global.DB.Delete(&advertList)
	response.OkWithMessage(fmt.Sprintf("一共删除 %d 个广告", count), c)
}
