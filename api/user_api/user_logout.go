package user_api

import (
	"gin_blog/global"
	response "gin_blog/models/common"
	"gin_blog/services"
	"gin_blog/utils/jwts"
	"github.com/gin-gonic/gin"
)

// LogoutView 用户注销
// @Tags 用户管理
// @Summary 用户注销
// @Description 用户注销
// @Param token header string  true  "token"
// @Router /api/logout [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) LogoutView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	token := c.Request.Header.Get("token")

	//{
	//	diff := claims.ExpiresAt.Time.Sub(time.Now())
	//	err := global.Redis.Set(fmt.Sprintf("logout_%s", token), "", diff).Err()
	//	if err != nil {
	//		response.FailWithMessage("注销失败", c)
	//		return
	//	}
	//}

	err := services.ServiceApp.UserService.Logout(claims, token)

	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("注销失败", c)
		return
	}

	response.OkWithMessage("注销成功", c)

}
