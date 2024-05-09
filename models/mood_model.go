package models

import "gin_blog/models/array_type"

type MoodModel struct {
	MODEL
	UserID  uint             `json:"user_id"`
	User    AuthModel        `gorm:"foreignKey:UserID" json:"-"` // 用户
	Avatar  string           `json:"avatar"`                     // 头像
	IP      string           `gorm:"size:32" json:"ip"`          // ip
	Addr    string           `gorm:"size:64" json:"addr"`
	Content string           `gorm:"size:256" json:"content"`
	Drawing array_type.Array `gorm:"type:longtext" json:"drawing"`
}
