package user_api

import (
	"gin_blog/global"
	"gin_blog/models"
	response "gin_blog/models/common"
	"gin_blog/utils/jwts"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"strings"
)

type UserUpdateNicknameRequest struct {
	NickName string `json:"nick_name" structs:"nick_name"`
	Sign     string `json:"sign" structs:"sign"`
	Link     string `json:"link" structs:"link"`
}

// UserUpdateNickName 修改当前登录人的昵称，签名，链接
// @Tags 用户管理
// @Summary 修改当前登录人的昵称，签名，链接
// @Description 修改当前登录人的昵称，签名，链接
// @Router /api/user_info [put]
// @Param token header string  true  "token"
// @Param data body UserUpdateNicknameRequest  true  "昵称，签名，链接"
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) UserUpdateNickName(c *gin.Context) {
	var cr UserUpdateNicknameRequest
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}
	var newMaps = map[string]interface{}{}
	maps := structs.Map(cr)
	for key, v := range maps {
		if val, ok := v.(string); ok && strings.TrimSpace(val) != "" {
			newMaps[key] = val
		}
	}
	var userModel models.UserModel
	err = global.DB.Debug().Take(&userModel, claims.UserID).Error
	if err != nil {
		response.FailWithMessage("用户不存在", c)
		return
	}
	err = global.DB.Model(&userModel).Updates(newMaps).Error
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("修改用户信息失败", c)
		return
	}
	response.OkWithMessage("修改个人信息成功", c)

}
