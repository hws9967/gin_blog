package advert_api

import (
	"gin_blog/global"
	"gin_blog/models"
	response "gin_blog/models/common"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

// AdvertUpdateView 更新广告
// @Tags 广告管理
// @Summary 更新广告
// @Description 更新广告
// @Param data body AdvertRequest    true  "广告的一些参数"
// @Param token header string  true  "token"
// @Param id path int  true  "id"
// @Router /api/adverts/{id} [put]
// @Produce json
// @Success 200 {object} res.Response{}
func (AdvertApi) AdvertUpdateView(c *gin.Context) {
	id := c.Param("id")
	var cr AdvertRequest

	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}

	var advert models.AdvertModel
	err = global.DB.Take(&advert, id).Error

	if err != nil {
		response.FailWithMessage("广告不存在", c)
		return
	}
	//此时结构体还不能直接传数据
	maps := structs.Map(&cr)
	err = global.DB.Model(&advert).Updates(maps).Error
	if err != nil {
		response.FailWithMessage("修改广告失败", c)
		return

		response.OkWithMessage("修改广告成功", c)
	}

}
