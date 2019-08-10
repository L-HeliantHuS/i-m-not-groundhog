package service

import (
	"go-crud/model"
	"go-crud/serializer"
)

// UserInfoService 用户普通信息
type UserInfoService struct {

}

// info 显示用户普通信息
func (service *UserInfoService) Info(id string) serializer.Response {
	var user model.User
	err := model.DB.First(&user, id).Error
	if err != nil {
		return serializer.Response{
			Status: 40002,
			Msg:    "没有这个数据",
			Error: 	err.Error(),
		}
	}
	return serializer.Response{
		Data:   serializer.BuildInfo(user),
	}
}
