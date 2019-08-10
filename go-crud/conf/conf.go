package conf

import (
	"go-crud/cache"
	"go-crud/model"
	"go-crud/task"
	"os"

	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	godotenv.Load()

	// 读取翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		panic(err)
	}

	task.CronJob()

	// 连接数据库
	model.Database(os.Getenv("MYSQL_DSN"))
	cache.Redis()
	if os.Getenv("LIVEOBS_ADDR") == "" {
		panic("请配置RTMP地址")
	}
	if os.Getenv("LIVEFLV_ADDR") == "" {
		panic("请配置FLV地址")
	}
}
