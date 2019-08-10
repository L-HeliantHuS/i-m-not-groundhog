package service

import (
	"go-crud/model"
	"go-crud/serializer"
)

// DeleteVideoService 删除视频
type DeleteVideoService struct {
}

// delete 删除此视频投稿
func (service *DeleteVideoService) Delete(id string) serializer.Response {
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Status: 40002,
			Msg:    "没有这个数据",
			Error: 	err.Error(),
		}
	}

	err = model.DB.Delete(&video).Error
	if err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库删除失败",
			Error: 	err.Error(),
		}
	}

	return serializer.Response{
		Data:   serializer.BuildVideo(video),
	}
}
