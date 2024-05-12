package lm

import (
	"gin_blog/global"
	"gin_blog/plugins/logstash/log_model"
	"github.com/gin-gonic/gin"
)

type loggerFace interface {
	Send(msg string, c *gin.Context)
	SendDebug(msg string, c *gin.Context)
	SendInfo(msg string, c *gin.Context)
	SendError(msg string, c *gin.Context)
	SendWarn(msg string, c *gin.Context)
}

func (a logger) Send(msg string, c *gin.Context) {
	_, err := a.save(log_model.InfoLevel, msg, c)
	if err != nil {
		global.LOG.Error(err.Error())
	}
}

func (a logger) SendDebug(msg string, c *gin.Context) {
	a.save(log_model.DebugLevel, msg, c)
}
func (a logger) SendInfo(msg string, c *gin.Context) {
	a.save(log_model.InfoLevel, msg, c)
}
func (a logger) SendWarn(msg string, c *gin.Context) {
	a.save(log_model.WarnLevel, msg, c)
}
func (a logger) SendError(msg string, c *gin.Context) {
	a.save(log_model.ErrorLevel, msg, c)
}

func (a logger) save(level log_model.Level, msg string, c *gin.Context) (log_model.LogModel, error) {
	ip := c.ClientIP()
	user := c.GetString("user")
	//addr := c.Request.Host
	addr := c.Request.URL.Path
	var userID uint

	logModel := log_model.LogModel{
		IP:       ip,
		NickName: user,
		Msg:      msg,
		Level:    level,
		Addr:     addr,
		UserID:   userID,
	}
	err := global.DB.Create(&logModel).Error
	return logModel, err
}
