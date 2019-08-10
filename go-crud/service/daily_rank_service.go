package service

import (
	"fmt"
	"go-crud/cache"
	"go-crud/model"
	"go-crud/serializer"
	"strings"
)

// DailyRankService 每日排行
type DailyRankService struct {
}

// rank 获取排行榜
func (service *DailyRankService) Rank() serializer.Response {
	var video []model.Video

	// 从Redis获取点击前10的视频
	vids, _ := cache.RedisClient.ZRevRange(cache.DailyRankKey, 0, 9).Result()

	if len(vids) > 0 {
		order := fmt.Sprintf("FIELD(id, %s)", strings.Join(vids, ","))
		err := model.DB.Where("id IN (?)", vids).Order(order).Find(&video).Error
		if err != nil {
			return serializer.Response{
				Status: 40002,
				Msg:    "没有这个数据",
				Error: 	err.Error(),
			}
		}
	}
	return serializer.Response{
		Data:   serializer.BuildVideos(video),
	}
}
