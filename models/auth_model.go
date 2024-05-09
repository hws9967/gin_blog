package models

import "gin_blog/models/status_type"

const (
	PermissionAdmin       = 1 // 管理员
	PermissionUser        = 2 // 普通登录人
	PermissionVisitor     = 3 // 游客
	PermissionDisableUser = 4 // 被禁用的用户
)

// AuthModel 用户表
type AuthModel struct {
	MODEL
	NickName       string                        `gorm:"size:42" json:"nick_name"`                                                         // 昵称
	UserName       string                        `gorm:"size:36" json:"user_name"`                                                         // 用户名
	Password       string                        `gorm:"size:64" json:"password"`                                                          // 密码
	AvatarID       uint                          `json:"avatar_id"`                                                                        // 头像id
	Avatar         ImageModel                    `json:"-"`                                                                                // 头像
	Email          string                        `json:"email"`                                                                            // 邮箱
	Tel            string                        `gorm:"size:18" json:"tel"`                                                               // 手机号
	Addr           string                        `gorm:"size:64" json:"addr"`                                                              // 地址
	Token          string                        `gorm:"size:64" json:"token"`                                                             // 其他平台的唯一id
	IP             string                        `gorm:"size:20" json:"IP"`                                                                //  ip地址
	Role           int                           `gorm:"size:4;default:1" json:"role"`                                                     // 权限  1 管理员  2 普通用户  3 游客
	SignStatus     status_type.AccountStatusType `gorm:"type=smallint(6)" json:"sign_status"`                                              // 注册来源
	ArticleModels  []ArticleModel                `gorm:"foreignKey:AuthID" json:"-"`                                                       // 发布的文章列表
	CollectsModels []ArticleModel                `gorm:"many2many:auth2_collects;joinForeignKey:AuthID;JoinReferences:ArticleID" json:"-"` // 收藏的文章列表
	SiteModels     []SiteModel                   `gorm:"many2many:auth_sites" json:"-"`                                                    // 收藏的网站列表
}

func ParseRole(role int) string {
	switch role {
	case PermissionAdmin:
		return "管理员"
	case PermissionUser:
		return "用户"
	case PermissionVisitor:
		return "游客"
	case PermissionDisableUser:
		return "被禁言"
	}
	return "其他"
}
