package service

import (
	"go-crud/model"
	"go-crud/serializer"
)

// ShowVideoService 视频详细
type ShowVideoService struct {

}

// show 显示视频详细
func (service *ShowVideoService) Show(id string) serializer.Response {
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Status: 40002,
			Msg:    "没有这个数据",
			Error: 	err.Error(),
		}
	}
	video.AddView()

	return serializer.Response{
		Data:   serializer.BuildVideo(video),
	}
}
