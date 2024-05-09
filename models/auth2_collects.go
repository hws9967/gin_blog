package models

import "time"

// Auth2Collects 自定义第三张表
type Auth2Collects struct {
	AuthID       uint         `gorm:"primaryKey"`
	AuthModel    AuthModel    `gorm:"foreignKey:AuthID"`
	ArticleID    uint         `gorm:"primaryKey"`
	ArticleModel ArticleModel `gorm:"foreignKey:ArticleID"`
	CreatedAt    time.Time
}
