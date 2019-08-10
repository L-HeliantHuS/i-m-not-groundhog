package model

import (
	"github.com/jinzhu/gorm"
	"go-crud/cache"
	"strconv"
)

type Video struct {
	gorm.Model
	Title  string
	Info   string
	URL    string
	Author uint
}

// View 点击数
func (video *Video) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.VideoViewKey(video.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}


// AddView点击数
func (video *Video) AddView()  {
	// 增加视频点击  view:video:10 点一次+1
	cache.RedisClient.Incr(cache.VideoViewKey(video.ID))

	// 增加每日排行榜点击数
	cache.RedisClient.ZIncrBy(cache.DailyRankKey, 1, strconv.Itoa(int(video.ID)))
	// 增加每周排行榜点击数
	cache.RedisClient.ZIncrBy(cache.WeeklyRankKey, 1, strconv.Itoa(int(video.ID)))
}
