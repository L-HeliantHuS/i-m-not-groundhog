package cache

import (
	"fmt"
	"strconv"
)

const (
	// 每日排行
	DailyRankKey = "rank:daily"

	// 每周排行
	WeeklyRankKey = "rank:weekly"
)

// VideoViewKey  视频点击数的Key
func VideoViewKey(id uint) string {
	// 返回view:video:id 格式的键名
	return fmt.Sprintf("view:video:%s", strconv.Itoa(int(id)))
}

