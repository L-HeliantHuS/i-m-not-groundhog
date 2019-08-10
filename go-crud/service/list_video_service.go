package service

import (
	"go-crud/model"
	"go-crud/serializer"
)

type ListVideoService struct {

}

// list 显示视频列表
func (service *ListVideoService) List() serializer.Response {
	var videos []model.Video
	err := model.DB.Find(&videos).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "数据库操作异常",
			Error: 	err.Error(),
		}
	}

	return serializer.Response{
		Data:   serializer.BuildVideos(videos),
	}
}
