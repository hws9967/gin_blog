package user_api

import (
	"gin_blog/models"
	response "gin_blog/models/common"
	"gin_blog/models/ctype"
	"gin_blog/services/common"
	"gin_blog/utils/desens"
	"gin_blog/utils/jwts"
	"github.com/gin-gonic/gin"
)

type UserResponse struct {
	models.UserModel
	RoleID int `json:"role_id"`
}

type UserListRequest struct {
	models.PageInfo
	Role int `json:"role" form:"role"`
}

// UserListView 用户列表
// @Tags 用户管理
// @Summary 用户列表
// @Description 用户列表
// @Param data query models.PageInfo  false  "查询参数"
// @Param token header string  true  "token"
// @Router /api/users [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.UserModel]}
func (UserApi) UserListView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	/*token := c.Request.Header.Get("token")
	if token == "" {
		response.FailWithMessage("未携带token", c)
		return
	}
	claims, err := jwts.ParseToken(token)
	if err != nil {
		response.FailWithMessage("token解析错误", c)
		return
	}*/
	var page UserListRequest
	if err := c.ShouldBindQuery(&page); err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}
	var users []UserResponse
	list, count, _ := common.ComList(models.UserModel{}, common.Option{
		PageInfo: page.PageInfo,
		Likes:    []string{"nick_name"},
	})

	for _, user := range list {
		//不是管理员我就不给你提供用户名
		if ctype.Role(claims.Role) != ctype.PermissionAdmin {
			user.UserName = ""
		}
		//电话和邮箱都脱敏处理
		user.Tel = desens.DesensitizationTel(user.Tel)
		user.Email = desens.DesensitizationEmail(user.Email)
		users = append(users, UserResponse{
			UserModel: user,
			RoleID:    int(user.Role),
		})
	}
	response.OkWithList(users, count, c)
}
