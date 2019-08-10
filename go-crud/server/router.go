package server

import (
	"go-crud/api"
	"go-crud/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		// 测试返回pong 和 hello World
		v1.POST("ping", api.Ping)


		// 用户注册登录
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)


		// 获取视频列表和查看单个视频   不需要登录
		v1.GET("video/:id", api.ShowVideo)
		v1.GET("videos", api.ListVideo)
		v1.GET("/user/user_info/:uid", api.UserInfo) // 通过uid获取用户不敏感信息


		// 排行榜
		v1.GET("/rank/daily", api.DailyRank)   // 每日
		v1.GET("/rank/weekly", api.WeeklyRank) // 每周


		// 需要登录保护的
		v1.Use(middleware.AuthRequired())
		{
			// User Routing
			v1.GET("user/me", api.UserMe)            // 通过cookie获取自己的敏感信息
			v1.DELETE("user/logout", api.UserLogout) // 注销用户

			// 投稿视频 修改视频 删除视频
			v1.POST("video", api.CreateVideo) // 投稿视频
			v1.PUT("videos/:id", api.UpdateVideo)
			v1.DELETE("videos/:id", api.DeleteVideo)
		}
	}
	return r
}
