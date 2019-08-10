package task

import "go-crud/cache"

func RestartDailyRank() error {
	return cache.RedisClient.Del(cache.DailyRankKey).Err()
}