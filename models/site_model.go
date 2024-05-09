package models

import "gin_blog/models/status_type"

// SiteModel 站点表
type SiteModel struct {
	MODEL
	Title         string                     `gorm:"size:32" json:"title"`                 // 标题
	Abstract      string                     `json:"abstract"`                             // 简介
	Url           string                     `json:"url"`                                  // 跳转地址
	Icon          string                     `json:"icon"`                                 // 网站图标
	Status        status_type.SiteStatusType `gorm:"type=int;default=0" json:"status"`     // 审核状态
	Msg           *string                    `gorm:"size:32" json:"msg"`                   // 错误描述
	Email         *string                    `gorm:"size:64" json:"email"`                 // 邮箱
	SiteType      int                        `gorm:"size:1;default:2" json:"site_type"`    // 网站的类型  1 博客 2 普通网站
	CollectsCount int                        `json:"collects_count"`                       // 收藏数
	DiggCount     int                        `json:"digg_count"`                           // 点赞数
	TagModels     []TagModel                 `gorm:"many2many:site_tag" json:"tag_models"` // 网站标签
}
