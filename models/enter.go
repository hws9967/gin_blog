package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primarykey" json:"id"` // 主键ID
	CreatedAt time.Time `json:"created_at"`           // 创建时间
	UpdatedAt time.Time `json:"-"`                    // 更新时间
}

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page  int      `json:"page" form:"page"`   // 页码
	Limit int      `json:"limit" form:"limit"` // 每页条数
	Key   string   `json:"key" form:"key"`     // 关键字
	Sort  string   `json:"sort" form:"sort"`   // 排序
	Like  []string `json:"like" form:"like"`   // 需要模糊匹配的列表
}

// ResponseList 公共的返回
type ResponseList struct {
	List  any `json:"list"`
	Count any `json:"count"`
}
