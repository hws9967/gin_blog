package message_api

import (
	"fmt"
	"gin_blog/utils/jwts"
	"github.com/gin-gonic/gin"
)

// MessageUserListByMeView 我与其他用户的聊天列表
// @Tags 消息管理
// @Summary 我与其他用户的聊天列表
// @Description 我与其他用户的聊天列表
// @Router /api/message_users/me [get]
// @Param token header string  true  "token"
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[MessageUserListResponse]}
func (m MessageApi) MessageUserListByMeView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	c.Request.URL.RawQuery = fmt.Sprintf("userID=%d", claims.UserID)
	m.MessageUserListByUserView(c)
}
