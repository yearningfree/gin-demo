package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

const (
	// InActiveUser 邮箱待激活用户
	InActiveUser string = "inactive"
	// ActiveUser 激活用户
	ActiveUser string = "active"
	// SuspendUser 封禁用户
	SuspendUser string = "suspend"
)

// User 用户模型
type User struct {
	gorm.Model
	UserName string `gorm:"not null"`         // 用户名
	PassWord string `gorm:"not null"`         // 密码
	NickName string `gorm:"type:varchar(20)"` // 昵称
	Status   string                           // 状态
	Avatar   string `gorm:"size:1000"`        // 头像
	Member   int    `gorm:"DEFAULT:0"`        // 会员身份，默认普通  会员等级1、2、3，待扩展
}

// GetUserByID 通过ID获取用户
func GetUserByID(ID interface{}) (User, error) {
	var user User
	result := DB.First(&user, ID)
	return user, result.Error
}

// EncryptedPassword 密码加密
func (user *User) EncryptedPassword(pwd string) error {
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), 12)
	if err != nil {
		return err
	}
	user.PassWord = string(hashPwd)
	return nil

}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(password))
	return err == nil
}
