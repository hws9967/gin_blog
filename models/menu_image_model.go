package models

// MenuImageModel 自定义菜单和背景图的连接表，方便排序
type MenuImageModel struct {
	MenuID     uint       `json:"menu_id"`
	MenuModel  MenuModel  `gorm:"foreignKey:MenuID"`
	ImageID    uint       `json:"image_id"`
	ImageModel ImageModel `gorm:"foreignKey:ImageID"`
	Sort       int        `gorm:"size:10" json:"sort"`
}
