package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primarykey" json:"id"` // 主键ID
	CreatedAt time.Time `json:"created_at"`           // 创建时间
	UpdatedAt time.Time `json:"-"`                    // 更新时间
}

// PageInfo 公共的请求参数
type PageInfo struct {
	Page  int    `form:"page"`  // 页码
	Key   string `form:"key"`   // 关键字
	Limit int    `form:"limit"` // 每页条数
	Sort  string `form:"sort"`  // 排序
}

type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}
