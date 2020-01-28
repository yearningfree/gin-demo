package serialzer

import "myownsite/server/model"

// User 用户序列化
type User struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"` // 用户名
	NickName  string `json:"nick_name"` // 昵称
	Status    string `json:"status"`    // 状态
	Avatar    string `json:"avatar"`    // 头像
	Member    int    `json:"member"`    // 会员身份，默认普通
	CreatedAt int64  `json:"created_at"`
}

// UserResponse 单个用户序列化
type UserResponse struct {
	Response
	Data User `json:"data"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:        user.ID,
		UserName:  user.UserName,
		NickName:  user.NickName,
		Status:    user.Status,
		Avatar:    user.Avatar,
		Member:    user.Member,
		CreatedAt: user.CreatedAt.Unix(),
	}

}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User) UserResponse {
	return UserResponse{
		Data:     BuildUser(user),
	}
}
