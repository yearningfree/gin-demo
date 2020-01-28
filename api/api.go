package api

import (
	"github.com/gin-gonic/gin"
	"gin-demo/model"
	"gin-demo/serialzer"
)

//Ping 检查
func Ping(c *gin.Context) {
	c.JSON(200, serialzer.Response{
		Status: 0,
		Msg:    "Pong",
	})
}

// GetUser 获取当前用户
func GetUser(c *gin.Context) *model.User {
	if user, _ := c.Get("user"); user != nil {
		if exist, ok := user.(*model.User); ok {
			return exist
		}
	}
	return nil
}

//ErrorResponse
func ErrorResponse(err error) serialzer.Response {
	return serialzer.Response{
		Status: 0,
		Error:  err.Error(),
	}
}
