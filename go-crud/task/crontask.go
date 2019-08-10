package task

import (
	"fmt"
	"github.com/robfig/cron"
	"reflect"
	"runtime"
	"time"
)

var Cron *cron.Cron

// 运行
func Run(job func() error)  {
	from := time.Now().UnixNano()
	err := job()
	to := time.Now().UnixNano()
	jobName := runtime.FuncForPC(reflect.ValueOf(job).Pointer()).Name()
	if err != nil {
		fmt.Printf("%s error: %dms\n", jobName, (to-from)/int64(time.Millisecond))
	} else {
		fmt.Printf("%s success: %dms\n", jobName,  (to-from)/int64(time.Millisecond))
	}
	
}


// 定时任务
func CronJob()  {
	if Cron == nil {
		Cron = cron.New()
	}

	// 每天情理redis排行榜
	Cron.AddFunc("0 0 1 * *", func() {
		Run(RestartDailyRank)
	})

	Cron.Start()

	fmt.Println("Cronjob starting...")

}