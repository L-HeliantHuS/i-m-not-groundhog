package service

import (
	"fmt"
	"go-crud/model"
	"go-crud/serializer"
)

// CreateVideoService 视频投稿服务
type CreateVideoService struct {
	Title  string `form:"title" json:"title" binding:"required,min=2,max=50"`
	Info   string `form:"info" json:"info" binding:"max=200"`
	URL    string `form:"url" json:"url" binding:"required,max=255"`
	UserID int    `form:"userid" json:"userid" binding:"required,max=255"`   // 千万不要写空格.
}

// create 创建视频
func (service *CreateVideoService) Create() serializer.Response {
	var user model.User
	user, err := model.GetUser(service.UserID)
	if err != nil {
		return serializer.Response{
			Status: 40000,
			Msg:    "我怀疑你在搞事情.",
			Error:  err.Error(),
		}
	}
	fmt.Println("username:" + user.Nickname)

	video := model.Video{
		Title: service.Title,
		Info:  service.Info,
		URL:   service.URL,
		Author: user.ID,
	}

	err = model.DB.Create(&video).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "视频保存失败",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}

}
