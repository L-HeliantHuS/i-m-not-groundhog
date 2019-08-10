package serializer

import (
	"go-crud/model"
)

// Video 视频序列化器   就是要返回json的数据
type Video struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Info      string `json:"info"`
	URL       string `json:"url"`
	Author    uint   `json:"author"`
	View      uint64 `json:"view"`
	CreatedAt int64  `json:"created_at"`
	UpdateAt  int64  `json:"updated_at"`
}

// 序列化单个视频
func BuildVideo(item model.Video) Video {
	return Video{
		ID:        item.ID,
		Title:     item.Title,
		Info:      item.Info,
		URL:       item.URL,
		Author:    item.Author,
		View:      item.View(),
		CreatedAt: item.CreatedAt.Unix(),
		UpdateAt:  item.UpdatedAt.Unix(),
	}
}

// 序列化视频列表
func BuildVideos(items []model.Video) (videos []Video) {
	for _, item := range items {
		video := BuildVideo(item)
		videos = append(videos, video)
	}
	return videos
}
