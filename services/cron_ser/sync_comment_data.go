package cron_ser

import (
	"gin_blog/global"
	"gin_blog/models"
	"gin_blog/services/redis_ser"
	"gorm.io/gorm"
)

// SyncCommentData 同步评论数据到数据库
func SyncCommentData() {
	commentDiggInfo := redis_ser.NewCommentDigg().GetInfo()
	for key, count := range commentDiggInfo {
		var comment models.CommentModel
		err := global.DB.Take(&comment, key).Error
		if err != nil {
			global.Log.Error(err)
			continue
		}
		err = global.DB.Model(&comment).
			Update("digg_count", gorm.Expr("digg_count + ?", count)).Error
		if err != nil {
			global.Log.Error(err)
			continue
		}
		global.Log.Infof("%s 更新成功 新的点赞数为：%d", comment.Content, comment.DiggCount)
	}
	redis_ser.NewCommentDigg().Clear()
}
