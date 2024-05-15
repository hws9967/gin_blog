package advert_api

import (
	"gin_blog/global"
	"gin_blog/models"
	response "gin_blog/models/common"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

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
		response.FailWithMessage("更新失败", c)
		return

		response.OkWithMessage("更新成功", c)
	}

}
