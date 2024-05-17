package menu_api

import (
	"gin_blog/global"
	"gin_blog/models"
	response "gin_blog/models/common"
	"github.com/gin-gonic/gin"
)

type MenuDetailRequest struct {
	Path string `json:"path" form:"path"`
}

func (MenuApi) MenuDetailByPathView(c *gin.Context) {
	// 先查菜单
	var cr MenuDetailRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}
	var menuModel models.MenuModel
	err = global.DB.Take(&menuModel, "path = ?", cr.Path).Error
	if err != nil {
		response.FailWithMessage("菜单不存在", c)
		return
	}
	// 查连接表
	var menuBanners []models.MenuBannerModel
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id = ?", menuModel.ID)
	var banners = make([]Banner, 0)
	for _, banner := range menuBanners {
		if menuModel.ID != banner.MenuID {
			continue
		}
		banners = append(banners, Banner{
			ID:   banner.BannerID,
			Path: banner.BannerModel.Path,
		})
	}
	menuResponse := MenuResponse{
		MenuModel: menuModel,
		Banners:   banners,
	}
	response.OkWithData(menuResponse, c)
	return
}
