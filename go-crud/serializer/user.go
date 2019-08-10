package serializer

import "go-crud/model"

// User 用户序列化器
type User struct {
	ID         uint   `json:"id"`
	UserName   string `json:"user_name"`
	Nickname   string `json:"nickname"`
	Signature  string `json:"signature"`   // 个性签名
	LiveAddr   string `json:"live_addr"`   // 直播间地址
	LiveStream string `json:"live_stream"` // 直播推流地址
	Status     string `json:"status"`
	Avatar     string `json:"avatar"`
	CreatedAt  int64  `json:"created_at"`
}

// UserResponse 单个用户序列化
type UserResponse struct {
	Response
	Data User `json:"data"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:         user.ID,
		UserName:   user.UserName,
		Nickname:   user.Nickname,
		Status:     user.Status,
		LiveStream: user.LiveStream,
		LiveAddr:   user.LiveAddr,
		CreatedAt:  user.CreatedAt.Unix(),
	}
}

func BuildInfo(user model.User) User {
	return User{
		ID:        user.ID,
		Nickname:  user.Nickname,
		LiveAddr:  user.LiveAddr,
		Status:    user.Status,
		CreatedAt: user.CreatedAt.Unix(),
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User) UserResponse {
	return UserResponse{
		Data: BuildUser(user),
	}
}
