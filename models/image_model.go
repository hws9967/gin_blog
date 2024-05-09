package models

import (
	"gin_blog/global"
	"gin_blog/models/image_type"
	"gorm.io/gorm"
	"os"
)

// ImageModel 图片表
type ImageModel struct {
	MODEL
	Path          string               `json:"path"`                                             // 图片路径
	ImageType     image_type.ImageType `gorm:"type=smallint(6)" json:"image_type"`               // 图片类型  封面 背景图 头像
	Sort          int                  `gorm:"size:4;default:0" json:"sort"`                     // 排序字段
	Key           string               `gorm:"column:image_key" json:"key"`                      // 图片的hash值，用于判断重复图片
	Name          string               `gorm:"size:38" json:"name"`                              // 图片名称
	NotUse        bool                 `gorm:"default:false;comment:是否可以选择，默认可以" json:"not_use"` // 是否可以选择
	ArticleModels []ArticleModel       `gorm:"foreignKey:CoverID" json:"-"`
	AuthModels    []AuthModel          `gorm:"foreignKey:AvatarID" json:"-"`
}

// BeforeDelete 删除的钩子函数
func (img *ImageModel) BeforeDelete(tx *gorm.DB) (err error) {
	// 将图片删除
	err = os.Remove(img.Path[1:])
	if err != nil {
		global.LOG.Error(err)
		return err
	}
	return nil
}
