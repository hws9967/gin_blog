package advert_api

import (
	"gin_blog/models"
	response "gin_blog/models/common"
	"gin_blog/services/common"
	"github.com/gin-gonic/gin"
	"strings"
)

// AdvertListView 广告列表
// @Tags 广告管理
// @Summary 广告列表
// @Description 广告列表
// @Param data query models.PageInfo    false  "查询参数"
// @Router /api/adverts [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.AdvertModel]}
func (AdvertApi) AdvertListView(c *gin.Context) {
	var cr models.PageInfo
	if err := c.ShouldBindQuery(&cr); err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}
	// 判断 Referer 是否包含admin，如果是，就全部返回，不是，就返回is_show=true
	referer := c.GetHeader("Referer")
	isShow := true
	if strings.Contains(referer, "admin") {
		// admin来的
		isShow = false
	}
	list, count, _ := common.ComList(models.AdvertModel{IsShow: isShow}, common.Option{
		PageInfo: cr,
		Debug:    true,
	})
	response.OkWithList(list, count, c)
}
