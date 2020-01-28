package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gin-demo/model"
	"gin-demo/serialzer"
)

// AuthenticationLogin 登录认证
func AuthorizedLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		uid := session.Get("user_id")
		if uid != nil {
			user, err := model.GetUserByID(uid)
			if err == nil {
				c.Set("user", &user)
			}
			c.Next()
		} else {

			c.JSON(200, serialzer.Response{
				Status: 5002, // 5002 服务器拒绝
				Msg:    "登录认证未通过，请重新登录",
			})
			c.Abort()
		}
	}
}
