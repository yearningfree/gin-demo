package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"myownsite/server/serialzer"
	"myownsite/server/services"
)

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	var userReg services.UserRegisterFrom
	if err := c.ShouldBind(&userReg); err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
	} else {
		if user, err := userReg.Register(); err != nil {
			c.JSON(200, err)
		} else {
			re := serialzer.BuildUserResponse(user)
			c.JSON(200, re)
		}
	}
}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var userLogin services.UserLoginFrom
	if err := c.ShouldBind(&userLogin); err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
	} else {
		if user, err := userLogin.Login(); err != nil {
			c.JSON(200, err)
		} else {
			s := sessions.Default(c)
			//s.Clear()
			s.Set("user_id", user.ID)
			s.Save()

			re := serialzer.BuildUserResponse(user)
			c.JSON(200, re)
		}
	}
}

// UserInfo 用户详情
func UserInfo(c *gin.Context) {
	user := GetUser(c)
	res := serialzer.BuildUserResponse(*user)
	c.JSON(200, res)

}

// UserLogout 用户登出
func UserLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(200, serialzer.Response{
		Status: 0,
		Msg:    "登出成功",
	})
}
