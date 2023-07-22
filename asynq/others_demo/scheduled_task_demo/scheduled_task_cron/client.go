package main

import (
	"fmt"
	"github.com/hibiken/asynq"
	"go-zero-micro/asynq/others_demo/async_task_demo/async_task_task"
	"log"
	"time"
)

const redisAddr = "127.0.0.1:6379"

func main() {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	scheduler := asynq.NewScheduler(
		&asynq.RedisClientOpt{
			Addr: redisAddr,
		},
		&asynq.SchedulerOpts{
			Location: loc,
		},
	)

	//emailBody := fmt.Sprintf("定时任务邮件已发送，发送时间：%s", time.Now().Format(utils.DateTimeFormat))
	emailBody := fmt.Sprintf("定时任务邮件已发送！")
	scheduledEmail := async_task_task.AsyncEmailPayload{
		To:      "user@example.com",
		Subject: "定时任务邮件",
		Body:    emailBody,
	}
	task, err := async_task_task.NewAsyncEmailTask(scheduledEmail)

	if err != nil {
		log.Fatal(err)
	}

	//entryID1, err := scheduler.Register("* * * * *", task) //每分钟执行一次任务
	//entryID1, err := scheduler.Register("*/1 * * * *", task) //每分钟执行一次任务
	entryID1, err := scheduler.Register("*/2 * * * *", task) //每2分钟执行一次任务
	//entryID1, err := scheduler.Register("@every 10s", task) //每隔10秒执行1次
	//entryID1, err := scheduler.Register("@every 1m", task)  //每隔1分钟执行1次
	//entryID1, err := scheduler.Register("@every 1h", task)  //每隔1小时执行1次
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("registered an entry: %q\n", entryID1)
	// 运行
	if err := scheduler.Run(); err != nil {
		log.Fatal(err)
	}
}
