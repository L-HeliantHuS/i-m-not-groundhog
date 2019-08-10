package service

import (
	"go-crud/model"
	"go-crud/serializer"
)

// UpdateVideoService 更新视频
type UpdateVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=50"`
	Info  string `form:"info" json:"info" binding:"max=200"`
}

// show 显示视频详细
func (service *UpdateVideoService) Update(id string) serializer.Response {
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Status: 40002,
			Msg:    "没有这个数据",
			Error: 	err.Error(),
		}
	}

	video.Title = service.Title
	video.Info = service.Info

	err = model.DB.Save(&video).Error
	if err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库保存失败",
			Error: 	err.Error(),
		}
	}

	return serializer.Response{
		Data:   serializer.BuildVideo(video),
	}
}
