package models

// SysModel 系统表
type SysModel struct {
	MODEL
	Version   string `gorm:"size:32" json:"version"`     // 网站版本
	SiteTitle string `gorm:"size:32" json:"site_title"`  // 网站标题
	SiteBeiAn string `gorm:"size:32" json:"site_bei_an"` // 网站备案号
}
