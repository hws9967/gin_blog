package user_api

import (
	"gin_blog/global"
	"gin_blog/models"
	response "gin_blog/models/common"
	"gin_blog/utils/jwts"
	"gin_blog/utils/pwd"
	"github.com/gin-gonic/gin"
)

type EmailLoginRequest struct {
	UserName string `json:"user_name" binding:"required" msg:"请输入用户名"`
	Password string `json:"password" binding:"required" msg:"请输入密码"`
}

// EmailLoginView 邮箱登录，返回token，用户信息需要从token中解码
// @Tags 用户管理
// @Summary 邮箱登录
// @Description 邮箱登录，返回token，用户信息需要从token中解码
// @Param data body EmailLoginRequest  true  "查询参数"
// @Router /api/email_login [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) EmailLoginView(c *gin.Context) {
	var cr EmailLoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}

	//log := log_stash.NewLogByGin(c)

	var userModel models.UserModel
	err = global.DB.Take(&userModel, "user_name = ? or email = ?", cr.UserName, cr.UserName).Error
	if err != nil {
		// 没找到
		global.Log.Warn("用户名不存在")
		//log.Warn(fmt.Sprintf("%s 用户名不存在", cr.UserName))
		response.FailWithMessage("用户名或密码错误", c)
		return
	}
	// 校验密码
	isCheck := pwd.CheckPwd(userModel.Password, cr.Password)
	if !isCheck {
		global.Log.Warn("用户名密码错误")
		//log.Warn(fmt.Sprintf("用户名密码错误 %s %s", cr.UserName, cr.Password))
		response.FailWithMessage("用户名或密码错误", c)
		return
	}
	// 登录成功，生成token
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		NickName: userModel.NickName,
		Role:     int(userModel.Role),
		UserID:   userModel.ID,
		Avatar:   userModel.Avatar,
	})
	if err != nil {
		global.Log.Error(err)
		//log.Error(fmt.Sprintf("token生成失败 %s", err.Error()))
		response.FailWithMessage("token生成失败", c)
		return
	}
	/*ip, addr := utils.GetAddrByGin(c)
	log = log_stash.New(ip, token)
	log.Info("登录成功")

	global.DB.Create(&models.LoginDataModel{
		UserID:    userModel.ID,
		IP:        ip,
		NickName:  userModel.NickName,
		Token:     token,
		Device:    "",
		Addr:      addr,
		LoginType: ctype.SignEmail,
	})*/

	response.OkWithData(token, c)

}
