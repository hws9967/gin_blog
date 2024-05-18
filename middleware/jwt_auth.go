package middleware

import (
	response "gin_blog/models/common"
	"gin_blog/models/ctype"
	"gin_blog/utils/jwts"
	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			response.FailWithMessage("未携带token", c)
			c.Abort()
			return
		}
		claims, err := jwts.ParseToken(token)
		if err != nil {
			response.FailWithMessage("token错误", c)
			c.Abort()
			return
		}
		//todo 用到redis再说
		// 判断是否在redis中
		/*if redis_ser.CheckLogout(token) {
			res.FailWithMessage("token已失效", c)
			c.Abort()
			return
		}*/
		// 登录的用户
		c.Set("claims", claims)
	}
}

func JwtAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			response.FailWithMessage("未携带token", c)
			c.Abort()
			return
		}
		claims, err := jwts.ParseToken(token)
		if err != nil {
			response.FailWithMessage("token错误", c)
			c.Abort()
			return
		}
		//todo 同上
		// 判断是否在redis中
		/*if redis_ser.CheckLogout(token) {
			res.FailWithMessage("token已失效", c)
			c.Abort()
			return
		}*/
		// 登录的用户
		if claims.Role != int(ctype.PermissionAdmin) {
			response.FailWithMessage("权限错误", c)
			c.Abort()
			return
		}

		c.Set("claims", claims)
	}
}
