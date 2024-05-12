package log_model

import "gin_blog/models"

type LogModel struct {
	models.MODEL
	IP        string           `gorm:"size:32" json:"ip"`                      // ip
	Addr      string           `gorm:"size:64" json:"addr"`                    // 用户地址
	UserID    uint             `json:"user_id"`                                // 用户id
	UserModel models.AuthModel `gorm:"foreignKey:UserID" json:"-"`             // 用户user
	NickName  string           `gorm:"size:16" json:"nick_name"`               // 操作的用户
	Level     Level            `gorm:"type:int;size:4;default:5" json:"level"` // 日志等级
	Msg       string           `gorm:"size:256" json:"msg"`                    // 消息
}
