// UserLoginService 管理用户注册的服务
package services

import (
	"myownsite/server/model"
	"myownsite/server/serialzer"
)

// UserRegisterFrom 用户注册数据表单
type UserRegisterFrom struct {
	Nickname        string `form:"nick_name" json:"nick_name" binding:"required,min=2,max=30"`
	UserName        string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password        string `form:"password" json:"pass_word" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
}

// Register 用户注册
func (userFrom *UserRegisterFrom) Register() (model.User, *serialzer.Response) {
	// 初始化用户信息
	user := model.User{
		UserName: userFrom.UserName,
		NickName: userFrom.Nickname,
		Status:   model.InActiveUser,
		Member:   0,
	}

	// 表单验证
	if err := userFrom.RegisterVaild(); err != nil {
		return user, err
	}

	// 加密密码
	err := user.EncryptedPassword(userFrom.Password)
	if err != nil {
		return user, &serialzer.Response{
			Status: 5002,
			Msg:    "密码加密失败",
		}
	}

	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return user, &serialzer.Response{
			Status: 5002,
			Msg:    "用户注册失败",
		}
	}

	return user, nil
}


// RegisterVaild 注册验证
func (userFrom *UserRegisterFrom) RegisterVaild() *serialzer.Response {
	if userFrom.Password != userFrom.PasswordConfirm {
		return &serialzer.Response{
			Status: 4002,
			Msg:    "两次输入的密码不相同！",
		}
	}

	userExist := 0
	model.DB.Model(&model.User{}).Where("user_name = ?", userFrom.UserName).Count(&userExist)
	if userExist > 0 {
		return &serialzer.Response{
			Status: 5001,
			Msg:    "用户名已被注册！",
		}
	}

	userExist = 0
	model.DB.Model(&model.User{}).Where("nick_name = ?", userFrom.Nickname).Count(&userExist)
	if userExist > 0 {
		return &serialzer.Response{
			Status: 5001,
			Msg:    "昵称已被占用！",
		}
	}

	return nil
}
