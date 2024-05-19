package user_api

import (
	"fmt"
	"gin_blog/global"
	"gin_blog/models"
	response "gin_blog/models/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserRemoveView 删除用户
// @Tags 用户管理
// @Summary 删除用户
// @Description 删除用户
// @Param data body models.RemoveRequest  true  "查询参数"
// @Param token header string  true  "token"
// @Router /api/users [delete]
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) UserRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	var userList []models.UserModel
	count := global.DB.Find(&userList, cr.IDList).RowsAffected
	if count == 0 {
		response.FailWithMessage("用户不存在", c)
		return
	}

	// 事务
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		// TODO:删除用户，消息表，评论表，用户收藏的文章，用户发布的文章
		err = global.DB.Delete(&userList).Error
		if err != nil {
			global.Log.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("删除用户失败", c)
		return
	}
	response.OkWithMessage(fmt.Sprintf("共删除 %d 个用户", count), c)

}
