package menu_api

import (
	"gin_blog/global"
	"gin_blog/models"
	response "gin_blog/models/common"
	"github.com/gin-gonic/gin"
)

type Banner struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
}

type MenuResponse struct {
	models.MenuModel
	Banners []Banner `json:"banners"`
}

func (MenuApi) MenuListView(c *gin.Context) {
	var menuList []models.MenuModel
	var menuIDList []uint
	//我先拿出menu的id
	global.DB.Order("sort desc").Find(&menuList).Select("id").Scan(&menuIDList)
	//再查询连接表
	var menuBanners []models.MenuBannerModel
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id in ?", menuIDList)
	var menus = make([]MenuResponse, 0)

	for _, menu := range menuList {
		var banners = make([]Banner, 0)
		for _, menuBanner := range menuBanners {
			if menuBanner.MenuID == menu.ID {
				banners = append(banners, Banner{
					ID:   menuBanner.BannerID,
					Path: menuBanner.BannerModel.Path,
				})
			}
		}
		menus = append(menus, MenuResponse{
			MenuModel: menu,
			Banners:   banners,
		})
	}

	response.OkWithList(menus, int64(len(menus)), c)
	return
}
