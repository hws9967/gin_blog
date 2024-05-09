package cmd

import (
	"gin_blog/global"
	"gin_blog/models"
)

// Makemigrations 初始化数据库表，数据库迁移等工作
func (CmderModel) Makemigrations() {
	var err error
	global.DB.SetupJoinTable(&models.AuthModel{}, "CollectsModels", &models.Auth2Collects{})
	global.DB.SetupJoinTable(&models.MenuModel{}, "MenuImages", &models.MenuImageModel{})
	// 生成四张表的表结构
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.ImageModel{},
			&models.TagModel{},
			&models.MessageModel{},
			&models.AdvertModel{},
			&models.AuthModel{},
			&models.CommentModel{},
			&models.ArticleModel{},
			&models.MenuModel{},
			&models.SiteModel{},
			&models.SysModel{},
			&models.MoodModel{},
			&models.Auth2Collects{},
			&models.MenuImageModel{},
			&models.FadeBackModel{},
			&models.LoginDataModel{},
		)
	if err != nil {
		global.LOG.Error("[ error ] 生成数据库表结构失败")
		return
	}
	global.LOG.Info("[ success ] 生成数据库表结构成功！")
}
