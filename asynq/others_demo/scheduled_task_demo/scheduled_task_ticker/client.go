package main

import (
	"fmt"
	"github.com/hibiken/asynq"
	"go-zero-micro/asynq/others_demo/async_task_demo/async_task_task"
	"go-zero-micro/common/utils"
	"log"
	"time"
)

const redisAddr = "127.0.0.1:6379"

func main() {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
	defer client.Close()

	StartUpScheduledTask(client, 10*time.Second) // 每隔 10 秒执行一次发送数据任务)
}

func StartUpScheduledTask(client *asynq.Client, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		emailBody := fmt.Sprintf("定时任务邮件已发送，发送时间：%s", time.Now().Format(utils.DateTimeFormat))
		scheduledEmail := async_task_task.AsyncEmailPayload{
			To:      "user@example.com",
			Subject: "定时任务邮件",
			Body:    emailBody,
		}
		task, err := async_task_task.NewAsyncEmailTask(scheduledEmail)
		//info, err := client.Enqueue(task)

		//※ asynq.ProcessAt(time.Now().Add(interval))是让服务端延迟指定的时间执行
		info, err := client.Enqueue(task, asynq.ProcessAt(time.Now().Add(interval)))

		if err != nil {
			log.Fatalf("could not enqueue task: %v", err)
		}
		log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
	}
}
