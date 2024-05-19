package user_api

import (
	"gin_blog/global"
	"gin_blog/models"
	response "gin_blog/models/common"
	"gin_blog/utils/jwts"
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
)

// UserInfoView 用户信息
// @Tags 用户管理
// @Summary 用户信息
// @Description 用户信息
// @Router /api/user_info [get]
// @Param token header string  true  "token"
// @Produce json
// @Success 200 {object} res.Response{data=models.UserModel}
func (UserApi) UserInfoView(c *gin.Context) {

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var userInfo models.UserModel
	err := global.DB.Take(&userInfo, claims.UserID).Error
	if err != nil {
		response.FailWithMessage("用户不存在", c)
		return
	}
	response.OkWithData(filter.Select("info", userInfo), c)

}
