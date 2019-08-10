package api

import (
	"github.com/gin-gonic/gin"
	"go-crud/service"
)

// rank/daily  每日排行榜接口
func DailyRank(c *gin.Context) {
	service := service.DailyRankService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Rank()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// rank/weekly  每周排行榜接口
func WeeklyRank(c *gin.Context) {
	service := service.WeeklyRankService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Rank()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}