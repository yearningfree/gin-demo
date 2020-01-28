// UserLoginService 管理用户登录的服务
package services

import (
	"myownsite/server/model"
	"myownsite/server/serialzer"
)

// UserLoginFrom 用户登录数据表单
type UserLoginFrom struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// Login 用户登录
func (userFrom *UserLoginFrom) Login() (model.User, *serialzer.Response) {
	var user model.User

	if err := model.DB.Where("user_name = ?", userFrom.UserName).First(&user).Error; err != nil {
		return user, &serialzer.Response{
			Status: 5001, // 服务器拒绝
			Msg:    "用户不存在",
		}
	}

	if user.CheckPassword(userFrom.Password) == false {
		return user, &serialzer.Response{
			Status: 4001, // 用户操作错误
			Msg:    "密码错误！",
		}
	}

	return user, nil
}
